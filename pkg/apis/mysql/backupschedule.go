// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupScheduleSpec defines the specification for a MySQL backup schedule.
type BackupScheduleSpec struct {
	// Schedule specifies the cron string used for backup scheduling.
	Schedule string

	// BackupTemplate is the specification of the backup structure
	// to get scheduled.
	BackupTemplate BackupSpec
}

// BackupSchedulePhase is a string representation of the lifecycle phase
// of a backup schedule.
type BackupSchedulePhase string

const (
	// BackupSchedulePhaseNew means the backup schedule has been created but not
	// yet processed by the backup schedule controller.
	BackupSchedulePhaseNew BackupSchedulePhase = "New"

	// BackupSchedulePhaseEnabled means the backup schedule has been validated and
	// will now be triggering backups according to the schedule spec.
	BackupSchedulePhaseEnabled BackupSchedulePhase = "Enabled"

	// BackupSchedulePhaseFailedValidation means the backup schedule has failed
	// the controller's validations and therefore will not trigger backups.
	BackupSchedulePhaseFailedValidation BackupSchedulePhase = "FailedValidation"
)

// ScheduleStatus captures the current state of a MySQL backup schedule.
type ScheduleStatus struct {
	// Phase is the current phase of the MySQL backup schedule.
	Phase BackupSchedulePhase

	// LastBackup is the last time a MySQLBackup was run for this
	// backup schedule.
	LastBackup metav1.Time
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackupScheduleList is a list of MySQLBackupSchedules.
type MySQLBackupScheduleList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []MySQLBackupSchedule
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackupSchedule is a MySQL Operator resource that represents a backup
// schedule of a MySQL cluster.
type MySQLBackupSchedule struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   BackupScheduleSpec
	Status ScheduleStatus
}