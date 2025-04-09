#### service port

| service name | api service port(1xxx) | rpc service port(2xxx) | prometheus(4xxx) |
| ------------ | ---------------------- | ---------------------- | ------------------------  |
| user         | 1001                   | 2001                   | 4001、4101                |
| asset        | 1002                   | 2002                   | 4002、4102                |
| operation    | 1003                   | 2003                   | 4003、4103                |
| workflow     | 1004                   | 2004                   | 4004、4104                |
| upload       | 1005                   | 2005                   | 4005、4105                |
| download     | 1006                   | -                      | 4006、-                   |
| monitor      | 1007                   | 2007                   | 4007、4107                |
| mqtt         | -                      | 2008                   | -、4108                   |
| robot        | 1009                   | 2009                   | 4009、4109                |