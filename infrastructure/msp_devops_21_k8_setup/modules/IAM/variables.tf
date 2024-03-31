variable "master-role-policy" {
  default = "gallery_policy_for_master_role"
}

variable "worker-role-policy" {
  default = "gallery_policy_for_worker_role"
}

variable "master-role" {
  default = "gallery_role_master_k8s"
}

variable "worker-role" {
  default = "gallery_role_worker_k8s"
}

variable "master-role-attachment" {
  default = "gallery_attachment_for_master"
}

variable "worker-role-attachment" {
  default = "gallery_attachment_for_worker"
}

variable "profile_for_master" {
  default = "gallery_profile_for_master"
}

variable "profile_for_worker" {
  default = "gallery_profile_for_worker"
}