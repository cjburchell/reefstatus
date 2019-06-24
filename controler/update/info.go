package update

import (
	"time"

	"github.com/cjburchell/reefstatus-common/data"
	"github.com/cjburchell/reefstatus-common/models"

	"github.com/cjburchell/profilux-go"
)

func info(controller *profilux.Controller, repo data.ControllerService) error {

	info, err := repo.GetInfo()
	if err != nil && err != data.ErrNotFound {
		return err
	}

	found := err != data.ErrNotFound

	info.LastUpdate = time.Now()

	info.SoftwareVersion, err = controller.GetSoftwareVersion()
	if err != nil {
		return err
	}

	info.Model, err = controller.GetModel()
	if err != nil {
		return err
	}

	info.SerialNumber, err = controller.GetSerialNumber()
	if err != nil {
		return err
	}

	info.SoftwareDate, err = controller.GetSoftwareDate()
	if err != nil {
		return err
	}

	info.DeviceAddress, err = controller.GetDeviceAddress()
	if err != nil {
		return err
	}

	info.Latitude, err = controller.GetLatitude()
	if err != nil {
		return err
	}

	info.Longitude, err = controller.GetLongitude()
	if err != nil {
		return err
	}

	info.MoonPhase, err = controller.GetMoonPhase()
	if err != nil {
		return err
	}
	info.Alarm, err = controller.GetAlarm()
	if err != nil {
		return err
	}

	info.OperationMode, err = controller.GetOperationMode()
	if err != nil {
		return err
	}

	for i := 0; i < 4; i++ {
		err = updateMaintenanceMode(controller, &info, i)
		if err != nil {
			return err
		}
	}

	reminderCount, err := controller.GetReminderCount()
	if err != nil {
		return err
	}

	for i := 0; i < reminderCount; i++ {
		err = updateReminder(controller, &info, i)
		if err != nil {
			return err
		}
	}

	return repo.SetInfo(info, !found)
}

func updateMaintenanceMode(controller *profilux.Controller, info *models.Info, index int) error {
	var maintenance *models.Maintenance
	for idx, item := range info.Maintenance {
		if item.Index == index {
			maintenance = &info.Maintenance[idx]
			break
		}
	}

	add := false
	if maintenance == nil {
		maintenance = models.NewMaintenance(index)
		add = true
	}

	var err error
	maintenance.Duration, err = controller.GetMaintenanceDuration(maintenance.Index)
	if err != nil {
		return err
	}
	maintenance.DisplayName, err = controller.GetMaintenanceText(maintenance.Index)
	if err != nil {
		return err
	}
	maintenance.IsActive, err = controller.IsMaintenanceActive(maintenance.Index)
	if err != nil {
		return err
	}
	maintenance.TimeLeft, err = controller.GetMaintenanceTimeLeft(maintenance.Index)
	if err != nil {
		return err
	}

	if add {
		info.Maintenance = append(info.Maintenance, *maintenance)
	}

	return nil
}

func updateReminder(controller *profilux.Controller, info *models.Info, index int) error {
	var reminder *models.Reminder
	var reminderIndex = 0
	for i, item := range info.Reminders {
		if item.Index == index {
			reminder = &info.Reminders[i]
			reminderIndex = i
			break
		}
	}

	isActive, err := controller.IsReminderActive(index)
	if err != nil {
		return err
	}

	if !isActive {
		if reminder != nil {
			info.Reminders = append(info.Reminders[:reminderIndex], info.Reminders[reminderIndex+1:]...)
		}
		return nil
	}

	add := false
	if reminder == nil {
		reminder = models.NewReminder(index)
		add = true
	}

	err = reminderUpdate(reminder, controller)
	if err != nil {
		return err
	}

	if add {
		info.Reminders = append(info.Reminders, *reminder)
	}

	return nil
}

// InfoState update
func InfoState(controller *profilux.Controller, repo data.ControllerService) error {
	info, err := repo.GetInfo()
	if err != nil {
		return err
	}

	info.LastUpdate = time.Now()
	info.Alarm, err = controller.GetAlarm()
	if err != nil {
		return err
	}
	info.OperationMode, err = controller.GetOperationMode()
	if err != nil {
		return err
	}
	info.MoonPhase, err = controller.GetMoonPhase()
	if err != nil {
		return err
	}

	for index := range info.Maintenance {
		info.Maintenance[index].IsActive, err = controller.IsMaintenanceActive(info.Maintenance[index].Index)
		if err != nil {
			return err
		}

		info.Maintenance[index].TimeLeft, err = controller.GetMaintenanceTimeLeft(info.Maintenance[index].Index)
		if err != nil {
			return err
		}
	}

	for _, item := range info.Reminders {
		err = reminderUpdate(&item, controller)
		if err != nil {
			return err
		}
	}

	return repo.SetInfo(info, false)
}

func reminderUpdate(reminder *models.Reminder, controller *profilux.Controller) error {
	err := reminderUpdateState(reminder, controller)
	if err != nil {
		return err
	}

	reminder.Text, err = controller.GetReminderText(reminder.Index)
	if err != nil {
		return err
	}

	reminder.Period, err = controller.GetReminderPeriod(reminder.Index)
	if err != nil {
		return err
	}

	reminder.IsRepeating, err = controller.GetReminderIsRepeating(reminder.Index)
	return err
}

func reminderUpdateState(reminder *models.Reminder, controller *profilux.Controller) error {
	var err error
	reminder.Next, err = controller.GetReminderNext(reminder.Index)
	if err != nil {
		return err
	}

	reminder.IsOverdue = reminder.Next.Before(time.Now())

	return nil
}
