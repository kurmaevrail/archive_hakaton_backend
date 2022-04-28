package main

type row struct {
    Time     float32 `ch:"Time"`
    Humidity float32 `ch:"Humidity"`
    TempRoom float32 `ch:"Room temperature"`
    TempZone float32 `ch:"Workspace temperature"`
    PH       float32 `ch:"pH"`
    Mass     float32 `ch:"Mass"`
    Waste    float32 `ch:"Waste"`
    CO2      float32 `ch:"CO2"`
}
