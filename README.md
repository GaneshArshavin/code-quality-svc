# code-quality-svc
## Description
The main aim of this service is to identify the type of files under a particular directory and the percentage of a particular language it holds within a repository.
### Config.json
The config.json file in the repository is used to indicate how to identify a particular file , it has 3 components 
```json
{
	"rules": [{
			"language": "go",
			"stratergy": "extension",
			"value": ".go"
		},
		{
			"language": "json",
			"stratergy": "extension",
			"value": ".json"
		},
		{
			"language": "makefile",
			"stratergy": "file_name",
			"value": "Makefile"
		},
		{
			"language": "Dockerfile",
			"stratergy": "file_name",
			"value": "Dockerfile"
		}
	]
}
```
language : specifies the language
stratergy : Stratergy enables the code to understand what to look for when its determining a particular file , this is because merely extensions are not enough for identifying
the type of file as complex files like Dockerfile and Makefile have no real extensions .
This also adds the scope to add more file discoverability in the future by adding a factory pattern .

The currently supported methods are detection are filename and extension .

## Execution

you can buid the CLI tool using the the command as follows  ```docker build . -t code-quality ``` . You can also run this on any specific directory by using the following command 
```docker run -it --volume {desired_directory}:/app  code-quality:latest findlanguage``` . Or you can atlernatively make use of the `run.sh` file as well . This would build the CLI and produce
the outputfile under the name ```findlanguageoutput```

The output is usually like 
```json
{
  "summary": {
    "Dockerfile": 0.0023255814,
    "go": 0.16511628,
    "makefile": 0.0023255814
  },
  "results": [
    {
      "path": "/app/Dockerfile",
      "language": "Dockerfile"
    },
    {
      "path": "/app/Feedback/cmd/client/main.go",
      "language": "go"
    },
    {
      "path": "/app/Feedback/cmd/server/main.go",
      "language": "go"
    },
    {
      "path": "/app/Feedback/pb/commons/commons.pb.go",
      "language": "go"
    },
    {
      "path": "/app/Feedback/pb/commons/tracing.pb.go",
      "language": "go"
    }
  ]
} 
```
The terms are self explanatory

##CI-CD Pipeline 
The CCID pipleline has two triggers , commits and tags . Unit tests ( a few sample tests has been added ) has been configured to for every commit. Tags trigger the deployment .
Unit tests are run recursively and is pre-requisite for builds to be pushed to docker hub 

 
             
 
          
