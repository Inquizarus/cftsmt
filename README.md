# CFTSMT

Small and simple tool to take a Terraform state output and present different kinds of information regarding its contents.

One of the main goals of this tool was to quickly generate remove/import statements for Terraform state manipulations and everything else is just a bonus.

## Usage

The most basic case would be to pipe the output of a Terraform show command into CFTSMT in JSON format.
```bash
$ terraform show -json | cftsmt <command> <flags>
```

So to generate a list of statements for all resources one would do like this.
```bash
$ terraform show -json | cftsmt resources statements
Remove                                                                  Import                                                                                                                                      
terraform state rm digitalocean_record.metrics                          terraform import digitalocean_record.metrics 12345678                                                          
terraform state rm digitalocean_record.webhooks                         terraform import digitalocean_record.webhooks 11345678                                                        
terraform state rm digitalocean_tag.backend                             terraform import digitalocean_tag.backend backend                                                               
terraform state rm digitalocean_tag.edge                                terraform import digitalocean_tag.edge edge                                                                     
terraform state rm digitalocean_tag.management                          terraform import digitalocean_tag.management management                                                         
terraform state rm digitalocean_tag.metrics                             terraform import digitalocean_tag.metrics metrics                                                               
terraform state rm digitalocean_tag.outbound                            terraform import digitalocean_tag.outbound outbound                                                             
terraform state rm digitalocean_tag.resources                           terraform import digitalocean_tag.resources resources                                                           
terraform state rm digitalocean_droplet.bastion                         terraform import digitalocean_droplet.bastion 872366344
```

To get statements for a specific resource one would just pass the flag `--resource-address` or for short `-a`.

```bash
$ terraform show -json | cftsmt resources statements --resource-address digitalocean_record.metrics
Remove                                          Import                                                 
terraform state rm digitalocean_record.metrics  terraform import digitalocean_record.metrics 12345678 
```

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

## TODO

- Add filtering on value/s when locating resources