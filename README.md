# Looker O2C Migration Evaluation

This Looker Open Source repository is released under the MIT license. This tool 
helps customer do self-evaluation of Looker Core migration. 
The tool is written in Go and uses the Looker SDK for Go.
There are several commands:
1. Compute Usage command to retrieve some looker usage data.
2. File System Performance command to assess the performance of the file system where Looker is hosted.

## Compute Usage
To use the script, you can run the following command from anywhere that can access the Looker instance:

```
go run main.go --command=compute-usage --api-id=<api_id> --api-secret=<api_secret> --addr=<looker_instance_address> --output-csv=<output_csv_path> --ssl=<ssl>
```

It will automatically output a CSV file with the usage details of the Looker
instance.

## File System Performance Usage
This script shows the size and file count of individual model-related directories, along with a disk write speed test.

Execute the following shell script on **your instance** only.
This script also assumes looker as the username and that the installation took place on the
user's home directory like we describe on our documentation https://cloud.google.com/looker/docs/installing-looker-application.

To use the script, you can run the following command:

```
go run main.go --command=file-system-performance --output-csv=<output_csv_path>
```

It will automatically output a CSV file with the file system performance of the Looker
instance.

