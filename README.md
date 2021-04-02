# learning-cloudformation
A repository to learn infrastructure automation with Cloudformation

TODO:
- [ ] Create a first version of the Jenkins stack
  - [ ] Use cross-stack references for the VPC creation. We don't need to update the VPC every time we deploy a new resource.
  - [ ] Define jenkins master and worker nodes with a single template
- [ ] Refactor the Jenkins-master-workers infrastructure with nested stacks
  - [ ] Create a folder structure in the repo to contain smaller templates that we can reuse (hint: we can call the directory `macros`)
  - [ ] Learn more about Lambda functions and how to use Macros effectively to create reusable cloudformation templates
  - [ ] Create a macro for ACM certificate generation
  - [ ] Create a macro for Jenkins master and workers security groups
  - [ ] Create a macro for Jenkins master instance and workers autoscaling group