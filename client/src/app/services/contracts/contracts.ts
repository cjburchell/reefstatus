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
