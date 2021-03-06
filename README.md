# CFTSMT

Small and simple tool to take a Terraform state output and present different kinds of information regarding its contents.

One of the main goals of this tool was to quickly generate remove/import statements for Terraform state manipulations and everything else is just a bonus.

## Usage

The most basic case would be to pipe the output of a Terraform show command into CFTSMT in JSON format.
```bash
$ terraform show -json | cftsmt <command> <flags>
```

## Statements list
Statements are defined as anything that manipulates a resource in the state.

To generate a list of statements for all resources one would do like this.
```bash
$ terraform show -json | cftsmt resources statements
Remove                                                                  Import                                                                                                                                  
terraform state rm digitalocean_record.metrics                          terraform import digitalocean_record.metrics  12345678                                                          
terraform state rm digitalocean_record.hooks                            terraform import digitalocean_record.hooks    22345678                                                        
terraform state rm digitalocean_tag.backend                             terraform import digitalocean_tag.backend     backend                                                               
terraform state rm digitalocean_tag.edge                                terraform import digitalocean_tag.edge        edge                                                                     
terraform state rm digitalocean_tag.management                          terraform import digitalocean_tag.management  management                                                         
terraform state rm digitalocean_tag.metrics                             terraform import digitalocean_tag.metrics     metrics                                                               
terraform state rm digitalocean_tag.outbound                            terraform import digitalocean_tag.outbound    outbound                                                             
terraform state rm digitalocean_tag.resources                           terraform import digitalocean_tag.resources   resources                                                           
terraform state rm digitalocean_droplet.bastion                         terraform import digitalocean_droplet.bastion 12345678
```

To get statements for a specific resource one would just pass the flag `--resource-address` or for short `-a`.

```bash
$ terraform show -json | cftsmt resources statements --resource-address digitalocean_record.metrics
Remove                                          Import                                                 
terraform state rm digitalocean_record.metrics  terraform import digitalocean_record.metrics 12345678 
```

## Filtering

Resource lists and statement lists can be filtered by three different kinds of filters in any combination.
There are `type`, `mode` and `values` filters.

With `type` you can specifiy something like `digitalocean_record` to find all resources of that type with the `--type-filter` flag.
With `mode` you can specify a Terraform mode like `data` or `managed` to find all resources of that mode with the `--mode-filter` flag.
With the `values` you can specify actual resource attributes as a valid JSON string to find all resources that matches those specific attributes and values with the `--values-filter` flag.

An example on the `values` filter below to get statements for all Digitalocean records with a TTL of 1800 seconds.

```bash
terraform show -json | cftsmt resources statements --values-filter '{"ttl":1800}'
Remove                                                                  Import                                                                         
terraform state rm digitalocean_record.auth                             terraform import digitalocean_record.auth     12345678                         
terraform state rm digitalocean_record.svc1                             terraform import digitalocean_record.svc1     22345678                        
terraform state rm digitalocean_record.metrics                          terraform import digitalocean_record.metrics  32345678           
terraform state rm digitalocean_record.hooks                            terraform import digitalocean_record.hooks    42345678 
```

## Other
For other types of usage just use the `-h` flag to get more information on available commands or flags for the same.
```bash
Usage:
  cftsmt [flags]
  cftsmt [command]

Available Commands:
  help        Help about any command
  modules     Lists all modules from the state in a neat table.
  outputs     Lists all outputs from the state in a neat table.
  resources   Lists all resources from the state in a neat table.

Flags:
  -h, --help   help for cftsmt
```
