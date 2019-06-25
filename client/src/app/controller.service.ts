import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface Maintenance {
  DisplayName: string;
  Index: number;
  IsActive: boolean;
  Duration: number;
  TimeLeft: number;
}

export interface Reminder {
  IsOverdue: boolean;
  Next: number;
  Text: string;
  Index: number;
  Period: number;
  IsRepeating: boolean;
}

export interface Info {
  OperationMode: string;
  Model: string;
  SoftwareDate: number;
  DeviceAddress: number;
  Latitude: number;
  Longitude: number;
  MoonPhase: number;
  Alarm: string;
  SoftwareVersion: number;
  SerialNumber: number;
  LastUpdate: number;
  Maintenance: Maintenance[];
  Reminders: Reminder[];
}

export interface LevelSensor {
  Id: string;
  DisplayName: string;
}

export interface SPort {
  Id: string;
  DisplayName: string;
}

export interface LPort {
  Id: string;
  DisplayName: string;
}

export interface DigitalInput {
  Id: string;
  DisplayName: string;
}

export interface Pump {
  Id: string;
  DisplayName: string;
}

export interface Light {
  Id: string;
  DisplayName: string;
}

export interface ProgrammableLogic {
  Id: string;
  DisplayName: string;
}


export interface DosingPump {
  Id: string;
  DisplayName: string;
}


export interface Probe  {
  Id: string;
  DisplayName: string;
  Units: string;
  Format: number;
  SensorType: string;
  Index: number;
  AlarmState: string;
  NominalValue: number;
  SensorMode: string;
  AlarmEnable: boolean;
  AlarmDeviation: number;
  Value: number;
  OperationHours: number;
  ConvertedValue: number;
  CenterValue: number;
  MaxRange: number;
  MinRange: number;
}

@Injectable({
  providedIn: 'root'
})
export class ControllerService {

  constructor(private http: HttpClient) { }

  getInfo() {
    return this.http.get<Info>('/api/v1/controller/info');
  }

  getProbes() {
    return this.http.get<Probe[]>('/api/v1/controller/probe');
  }

  getLevelSesnsors() {
    return this.http.get<LevelSensor[]>('/api/v1/controller/levelsensor');
  }

  getSPorts() {
    return this.http.get<SPort>('/api/v1/controller/sport');
  }

  getLPorts() {
    return this.http.get<LPort>('/api/v1/controller/lport');
  }

  getDigitalInputs() {
    return this.http.get<DigitalInput>('/api/v1/controller/digitalinput');
  }

  getPumps() {
    return this.http.get<Pump>('/api/v1/controller/pump');
  }

  getProgrammableLogic() {
    return this.http.get<ProgrammableLogic>('/api/v1/controller/programmablelogic');
  }

  getDosingPump() {
    return this.http.get<DosingPump>('/api/v1/controller/dosingpump');
  }

  getLight() {
    return this.http.get<Light>('/api/v1/controller/light');
  }
}
