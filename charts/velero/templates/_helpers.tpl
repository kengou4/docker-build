{*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*}

{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "velero.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "velero.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "velero.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create the name of the service account to use for creating or deleting the velero server
*/}}
{{- define "velero.serverServiceAccount" -}}
{{- if .Values.serviceAccount.server.create -}}
    {{ default (printf "%s-%s" (include "velero.fullname" .) "server") .Values.serviceAccount.server.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.server.name }}
{{- end -}}
{{- end -}}

{{/*
Create the name for the credentials secret.
*/}}
{{- define "velero.secretName" -}}
{{- if .Values.credentials.existingSecret -}}
  {{- .Values.credentials.existingSecret -}}
{{- else -}}
  {{- include "velero.fullname" . -}}
{{- end -}}
{{- end -}}

{{/*
Create the Velero priority class name.
*/}}
{{- define "velero.priorityClassName" -}}
{{- if .Values.priorityClassName -}}
  {{- .Values.priorityClassName -}}
{{- else -}}
  {{- include "velero.fullname" . -}}
{{- end -}}
{{- end -}}

{{/*
Create the Restic priority class name.
*/}}
{{- define "velero.restic.priorityClassName" -}}
{{- if .Values.restic.priorityClassName -}}
  {{- .Values.restic.priorityClassName -}}
{{- else -}}
  {{- include "velero.fullname" . -}}
{{- end -}}
{{- end -}}
