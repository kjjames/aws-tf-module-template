locals {
  enabled = var.enabled
  input = {
    enabled = var.enabled == null ? var.context.enabled : var.enabled
  }
}
