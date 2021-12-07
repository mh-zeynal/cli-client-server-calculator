# cli-client-server-calculator #
### a cli client-server app with cobra ###  
# overview #  
this project is a cli client-server app in which client gives a bunch of numbers with an operator to server  
and server does some calculations according to the operator and sends the result to client  
# commands #  
in order to see complete guide to commands of this code you'd better run `csp -h` or `csp --help`:  
``` 
Available Commands:  
  completion  generate the autocompletion script for the specified shell  
  enable      client-side starter command  
  help        Help about any command  
  server      runs server  
  
Flags:  
      --config string   config file (default is $HOME/.root.yaml)  
  -h, --help            help for root  
  -t, --toggle          Help message for toggle  
```  
`csp enable` and `csp server` are our major commands here:  
  - - - -
`csp enable`->  
to use this command we recommend you to use this pattern `csp enable FIRST_NUMBER SECOND_NUMBER ... nTH_NUMBER OPERATOR(+, -, *, /)` like so:  
![picture alt](/gifs/Record_2021_11_27_01_43_51_906_1.gif)  
NOTE: be careful that you must turn server on at first to get result  
  
  
`csp server`->  
to turn your server on you must use this command(in another CLI window) as below:  
![picture alt](/gifs/Record_2021_11_27_01_57_50_310.gif)  
  
__NOTE: this cli app has been tested inside jetbrains goland terminal__ 
