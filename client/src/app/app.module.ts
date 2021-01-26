import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import {InfoComponent} from './status/info.component';
import {LevelSensorComponent} from './status/levelSensor.component';
import {MaintenanceComponent} from './status/maintenance.component';
import {ProbeComponent} from './status/probe.component';
import {HttpClientModule} from '@angular/common/http';
import {AppRoutingModule} from './app-routing.module';

@NgModule({
  declarations: [
    AppComponent,
    InfoComponent,
    LevelSensorComponent,
    MaintenanceComponent,
    ProbeComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
