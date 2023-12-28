package run

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

func (r Run) Run() error {
	// Check if output files folder exists
	f, err := os.Stat(r.Config.OutputFiles)
	if f != nil && f.IsDir() {
		r.Logger.Error().Err(err).Str("output_files", r.Config.OutputFiles).Msg("output files folder already exists, you can't create a project in an existing folder")
		return err
	}

	if os.IsNotExist(err) {
		// Create output files folder
		err = os.MkdirAll(r.Config.OutputFiles, os.ModePerm)
		if err != nil {
			r.Logger.Error().Err(err).Str("output_files", r.Config.OutputFiles).Msg("failed to create output files folder")
			return err
		}
	}

	// Check if values file exists
	if _, err := os.Stat(r.Config.ValuesFile); os.IsNotExist(err) {
		r.Logger.Error().Err(err).Str("value_file", r.Config.ValuesFile).Msg("values file does not exist")
		return err
	}

	// Check if template files folder exists
	if _, err := os.Stat(r.Config.TemplateFiles); os.IsNotExist(err) {
		r.Logger.Error().Err(err).Str("template_files", r.Config.TemplateFiles).Msg("template files folder does not exist")
		return err
	}

	r.Logger.Info().Str("value_file", r.Config.ValuesFile).Msg("rendering values file")
	yFile, err := os.ReadFile(r.Config.ValuesFile)
	if err != nil {
		r.Logger.Error().Err(err).Msg("failed to read values file")
		return err
	}

	data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yFile, &data)
	if err != nil {
		r.Logger.Error().Err(err).Str("value_file", r.Config.ValuesFile).Msg("failed to unmarshal values file")
		return err
	}

	r.Logger.Info().Str("template_path", r.Config.TemplateFiles).Msg("parsing template files")
	filepath.Walk(r.Config.TemplateFiles, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			r.Logger.Error().Err(err).Str("template_file", path).Msg("failed to walk template file")
			return err
		}

		// Check if file is a directory
		if !info.IsDir() {
			// Parsez le template
			if !strings.Contains(path, "Templater.yaml") {
				tmpl, err := template.ParseFiles(path)
				if err != nil {
					r.Logger.Error().Err(err).Str("template_file", path).Msg("failed to parse template file")
					return err
				}

				// Construct the output file path
				outputFilePath := strings.Replace(path, r.Config.TemplateFiles, r.Config.OutputFiles, 1)

				// Remove the template file name from the output file path
				outputFilePath = strings.Replace(outputFilePath, info.Name(), "", 1)

				// Create the parent output directory if it doesn't exist
				if err := os.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm); err != nil {
					r.Logger.Error().Err(err).Str("output_file", outputFilePath).Msg("failed to create output file")
					return err
				}

				// Render the template file in memory
				r.Logger.Info().Str("output_file", outputFilePath).Msg("rendering template file")
				var renderedTemplate bytes.Buffer
				err = tmpl.Execute(&renderedTemplate, data)
				if err != nil {
					r.Logger.Error().Err(err).Str("output_file", outputFilePath).Msg("failed to execute template")
					return err
				}
				renderedStr := renderedTemplate.String()

				// Regex for finding the filename markers
				re := regexp.MustCompile(`(## filename: .+?\n)`)

				// Divide the rendered string into slices based on the filename markers
				splits := re.Split(renderedStr, -1)

				// Récupérer les noms de fichiers à partir des marqueurs
				fileNames := re.FindAllString(renderedStr, -1)

				if len(splits) <= 1 || len(splits) == 0 {
					// Create the output file
					outputFile, err := os.Create(outputFilePath + info.Name())
					if err != nil {
						r.Logger.Error().Err(err).Str("output_file", outputFilePath+info.Name()).Msg("failed to create output file")
						return err
					}
					defer outputFile.Close()

					// Write the rendered template to the output file
					r.Logger.Info().Str("output_file", outputFilePath+info.Name()).Msg("rendering template file")
					outputFile.WriteString(renderedStr)
				} else {
					for i, split := range splits {
						if i > 0 { // Ignore the first split
							filename := strings.TrimSpace(strings.TrimPrefix(fileNames[i-1], "## filename:"))
							content := strings.TrimSpace(split)

							r.Logger.Info().Str("filename", filename).Msg("rendering template file")

							// Create the output file
							outputFile, err := os.Create(outputFilePath + filename)
							if err != nil {
								r.Logger.Error().Err(err).Str("output_file", outputFilePath+filename).Msg("failed to create output file")
								return err
							}

							// Write the rendered template to the output file
							r.Logger.Info().Str("output_file", outputFilePath+filename).Msg("rendering template file")
							outputFile.WriteString("## filename: " + filename + "\n")
							outputFile.WriteString(content)
						}
					}
				}
			}
		}
		return nil
	})
	return nil
}
