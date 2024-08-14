# Helmdocs

The idea behind this generator is to create `readme` and template `values.yaml` files from the `values.schema.json` of your helm charts

Other generators use the `values.yaml` files to generate the rest. I think this approach is wrong because it is always best to start with strong typed definitions and then generate weaker typed files. In this case, values.yaml is constricted by object, basic types and hooks, while `values.schema.json` defines a.... schema.

I hope this is of some use and happy helming to everyone!