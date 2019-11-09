# AzureIoTHubClient
A go client to test the IoT hub custom routing

This project is part of a learning process to learn go  

To run the code execute the following command. 

```
go get -u github.com/amenzhinsky/iothub
go get -u github.com/tkanos/gonfig
```
In this example I choose to do the routing based upon a application property. 

In your IoT hub instance add a custom route e.g. DeviceStatusChanged in this case.

![custom route](/images/MessageRouting.JPG)


