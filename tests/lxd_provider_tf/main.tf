resource "lxd" "lxd_demo_container" {
	container_name = "alpine-c1"
	desired_status = "stop"
}