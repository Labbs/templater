package render

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/labbs/templater/bootstrap"
	"gopkg.in/yaml.v3"
)

func Render(app *bootstrap.Application) error {
	app.Logger.Info().Str("value_file", app.Config.ValuesFile).Msg("rendering values file")
	yFile, err := os.ReadFile(app.Config.ValuesFile)
	if err != nil {
		app.Logger.Error().Err(err).Msg("failed to read values file")
		return err
	}

	data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yFile, &data)
	if err != nil {
		app.Logger.Error().Err(err).Str("value_file", app.Config.ValuesFile).Msg("failed to unmarshal values file")
		return err
	}

	app.Logger.Info().Str("template_path", app.Config.TemplateFiles).Msg("parsing template files")
	filepath.Walk(app.Config.TemplateFiles, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			app.Logger.Error().Err(err).Str("template_file", path).Msg("failed to walk template file")
			return err
		}

		// Assurez-vous que le fichier n'est pas un dossier
		if !info.IsDir() {
			// Parsez le template
			if !strings.Contains(path, "Templater.yaml") {
				tmpl, err := template.ParseFiles(path)
				if err != nil {
					app.Logger.Error().Err(err).Str("template_file", path).Msg("failed to parse template file")
					return err
				}

				// Construisez le chemin de sortie
				outputFilePath := strings.Replace(path, app.Config.TemplateFiles, app.Config.OutputFiles, 1)

				// Supprimez le nom du fichier de sortie
				outputFilePath = strings.Replace(outputFilePath, info.Name(), "", 1)

				// Créez les dossiers parents au besoin
				if err := os.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm); err != nil {
					app.Logger.Error().Err(err).Str("output_file", outputFilePath).Msg("failed to create output file")
					return err
				}

				// Effectuez le rendu en mémoire
				app.Logger.Info().Str("output_file", outputFilePath).Msg("rendering template file")
				var renderedTemplate bytes.Buffer
				err = tmpl.Execute(&renderedTemplate, data)
				if err != nil {
					app.Logger.Error().Err(err).Str("output_file", outputFilePath).Msg("failed to execute template")
					return err
				}
				renderedStr := renderedTemplate.String()

				// Expression régulière pour trouver les marqueurs
				re := regexp.MustCompile(`(## filename: .+?\n)`)

				// Trouver tous les marqueurs et diviser le rendu
				splits := re.Split(renderedStr, -1)

				// Récupérer les noms de fichiers à partir des marqueurs
				fileNames := re.FindAllString(renderedStr, -1)

				if len(splits) <= 1 || len(splits) == 0 {
					// Créez le fichier de sortie
					outputFile, err := os.Create(outputFilePath + info.Name())
					if err != nil {
						app.Logger.Error().Err(err).Str("output_file", outputFilePath+info.Name()).Msg("failed to create output file")
						return err
					}
					defer outputFile.Close()

					// Effectuez le rendu dans le fichier de sortie
					app.Logger.Info().Str("output_file", outputFilePath+info.Name()).Msg("rendering template file")
					outputFile.WriteString(renderedStr)
				} else {
					for i, split := range splits {
						if i > 0 { // Ignorer le premier élément si c'est vide
							filename := strings.TrimSpace(strings.TrimPrefix(fileNames[i-1], "## filename:"))
							content := strings.TrimSpace(split)

							app.Logger.Info().Str("filename", filename).Msg("rendering template file")

							// Créez le fichier de sortie
							outputFile, err := os.Create(outputFilePath + filename)
							if err != nil {
								app.Logger.Error().Err(err).Str("output_file", outputFilePath+filename).Msg("failed to create output file")
								return err
							}

							// Effectuez le rendu dans le fichier de sortie
							app.Logger.Info().Str("output_file", outputFilePath+filename).Msg("rendering template file")
							outputFile.WriteString("## filename: " + filename + "\n")
							outputFile.WriteString(content)
						}
					}
				}

				// // Créez le fichier de sortie
				// outputFile, err := os.Create(outputFilePath)
				// if err != nil {
				// 	app.Logger.Error().Err(err).Str("output_file", outputFilePath).Msg("failed to create output file")
				// 	return err
				// }
				// defer outputFile.Close()

				// // Effectuez le rendu dans le fichier de sortie
				// app.Logger.Info().Str("output_file", outputFilePath).Msg("rendering template file")
				// err = tmpl.Execute(outputFile, data)
				// if err != nil {
				// 	app.Logger.Error().Err(err).Str("output_file", outputFilePath).Msg("failed to execute template")
				// 	return err
				// }
			}
		}
		return nil
	})
	return nil
}
