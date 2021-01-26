import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {
  DigitalInput,
  DosingPump,
  Info,
  LevelSensor,
  Light,
  LPort,
  Probe,
  ProgrammableLogic,
  Pump,
  SPort
} from './contracts/contracts';

@Injectable({
  providedIn: 'root'
})
export class ControllerService {

  constructor(private http: HttpClient) { }

  async getInfo(): Promise<Info> {
    return this.http.get<Info>('/api/v1/controller/info').toPromise();
  }

  getProbes(): Promise<Probe[]> {
    return this.http.get<Probe[]>('/api/v1/controller/probe').toPromise();
  }

  getLevelSesnsors(): Promise<LevelSensor[]> {
    return this.http.get<LevelSensor[]>('/api/v1/controller/levelsensor').toPromise();
  }

  getSPorts(): Promise<SPort[]> {
    return this.http.get<SPort[]>('/api/v1/controller/sport').toPromise();
  }

  getLPorts(): Promise<LPort[]> {
    return this.http.get<LPort[]>('/api/v1/controller/lport').toPromise();
  }

  getDigitalInputs(): Promise<DigitalInput[]> {
    return this.http.get<DigitalInput[]>('/api/v1/controller/digitalinput').toPromise();
  }

  getPumps(): Promise<Pump[]> {
    return this.http.get<Pump[]>('/api/v1/controller/pump').toPromise();
  }

  getProgrammableLogic(): Promise<ProgrammableLogic[]> {
    return this.http.get<ProgrammableLogic[]>('/api/v1/controller/programmablelogic').toPromise();
  }

  getDosingPump(): Promise<DosingPump[]> {
    return this.http.get<DosingPump[]>('/api/v1/controller/dosingpump').toPromise();
  }

  getLight(): Promise<Light[]> {
    return this.http.get<Light[]>('/api/v1/controller/light').toPromise();
  }
}
