import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface Maintenance{
  DisplayName: string,
  Index:       number,
  IsActive:    boolean,
  Duration:    number
  TimeLeft:    number
}

export interface Reminder{
  IsOverdue   :boolean,
  Next        :number,
  Text        :string,
  Index       :number,
  Period      :number,
  IsRepeating :boolean
}

export interface Info {
  OperationMode   :string
  Model           :string
  SoftwareDate    :number
  DeviceAddress   :number
  Latitude        :number
  Longitude       :number
  MoonPhase       :number
  Alarm           :string
  SoftwareVersion :number
  SerialNumber    :number
  LastUpdate      :number
  Maintenance: Maintenance[],
  Reminders: Reminder[]
}

export interface Probe  {
  Id             :string
  DisplayName    :string
  Units          :string
  Format         :number
  SensorType     :string
  Index          :number
  AlarmState     :string
  NominalValue   :number
  SensorMode     :string
  AlarmEnable    :boolean
  AlarmDeviation :number
  Value          :number
  OperationHours :number
  ConvertedValue :number
  CenterValue    :number
  MaxRange       :number
  MinRange       :number
}

@Injectable({
  providedIn: 'root'
})
export class ControllerService {

  constructor(private http: HttpClient) { }

  getInfo() {
    return this.http.get<Info>("/controller/info");
  }

  getProbes() {
    return this.http.get<Probe[]>("/controller/probe");
  }

  getLevelSesnsors() {
    return this.http.get<Info>("/controller/levelsensor");
  }

  getSPorts() {
    return this.http.get<Info>("/controller/sport");
  }

  getLPorts() {
    return this.http.get<Info>("/controller/lport");
  }

  getDigitalInputs() {
    return this.http.get<Info>("/controller/digitalinput");
  }

  getPumps() {
    return this.http.get<Info>("/controller/pump");
  }

  getProgrammableLogic() {
    return this.http.get<Info>("/controller/programmablelogic");
  }

  getDosingPump() {
    return this.http.get<Info>("/controller/dosingpump");
  }

  getLight() {
    return this.http.get<Info>("/controller/light");
  }
}
