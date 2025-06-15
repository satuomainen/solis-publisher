# solis-publisher

Get current power output from Solis API and publish it to an MQTT topic.

The Solis API is providing power station metrics from a Solis inverter
installed at a solar panel power station. The inverter is equipped with
a data logger that is collecting measurements from the inverter and storing
them in the Solis Cloud system. Data from the Solis Cloud can be read using
the Solis API with API keys associated to the registered owner of the data
logger.

```mermaid
C4Deployment
    Deployment_Node(controlService, "Controller", "My hobby server") {
        Container(updateCurrentPower, "update_current_power", "Cron Job")

        Rel(updateCurrentPower, solisCloud, "Fetch current power production and publish it to MQTT topic")
        Rel(updateCurrentPower, mqttTopic, "Publish current power production", "yieldkw")

        UpdateRelStyle(updateCurrentPower, mqttTopic, $offsetX="-45")
        UpdateRelStyle(updateCurrentPower, solisCloud, $offsetX="-45", $offsetY="15")
    }

    Deployment_Node(solisBoundary, "Solis", "Solis inverter HW and services") {
        Container(solisInverter, "Inverter", "Solar power inverter")
        Container(solisLogger, "Datalogger", "Serial device")
        Container(solisCloud, "Solis Cloud", "Cloud-based solar data service")

        Rel(solisLogger, solisInverter, "Collect data")
        Rel(solisLogger, solisCloud, "Send readings")
    }

    Deployment_Node(mqttCloud, "MQTT Broker Service", "") {
        Container(mqttTopic, "MQTT Topic", "yieldkw")
    }

    Deployment_Node(shellyRelay, "Shelly Pro 1", "Smart Relay") {
        Container(internalMqttClient, "Internal MQTT client")
        Container(runWhenSunny, "Script", "Control loads MQTT value")

        Rel(internalMqttClient, runWhenSunny , "MQTT values")
        Rel(internalMqttClient, mqttTopic, "Subscribe", "yieldkw")

        UpdateRelStyle(internalMqttClient, mqttTopic, $offsetX="-35", $offsetY="10")
        UpdateRelStyle(internalMqttClient, runWhenSunny, $offsetX="-40", $offsetY="-15")
    }
```

## Intended operation

This module provides a command line tool that can be periodically run for
example as a cron job. The update interval should probably not be shorter
than 5 minutes.

## Use cases

* Turn loads on/off based on the amount of power generated
* Graph the energy yield over time (Solis Cloud already does this)
