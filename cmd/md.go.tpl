{{- $title := .Title }}
{{- $desc := .Description}}
{{- range .Components}}
{{- if eq . "Title & Description"}}
# {{$title}}
{{$desc}}
{{- end}}

{{- if eq . "Demo Section"}}
## Demo

Insert gif or link to demo
{{- end}}

{{- if eq . "Installation"}}
## Installation

Install my-project with npm

```bash
  npm install my-project
  cd my-project
```
{{- end}}
{{- if eq . "Run locally"}}
## Run Locally

Clone the project

```bash
  git clone https://link-to-project
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  npm install
```

Start the server

```bash
  npm run start
```
{{- end}}
{{- if eq . "Env Variables"}}
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`API_KEY`

`ANOTHER_API_KEY`
{{- end}}
{{- if eq . "FAQ"}}
## FAQ

#### Question 1

Answer 1

#### Question 2

Answer 2
{{- end}}

{{- if eq . "ScreenShots"}}
## Screenshots

![App Screenshot](https://via.placeholder.com/468x300?text=App+Screenshot+Here)
{{- end}}

{{- if eq . "Acknowledgments"}}
## Acknowledgements

 - [Awesome Readme Templates](https://awesomeopensource.com/project/elangosundar/awesome-README-templates)
 - [Awesome README](https://github.com/matiassingers/awesome-readme)
 - [How to write a Good readme](https://bulldogjob.com/news/449-how-to-write-a-good-readme-for-your-github-project)

{{- end}}
{{- if eq . "License"}}
## License

[MIT](https://choosealicense.com/licenses/mit/)

{{- end}}
{{- end}}