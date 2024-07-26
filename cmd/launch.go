package cmd

import (
	"dockit/docker"
	"fmt"
)

type Launch struct{}

func (launch *Launch) MissingParams() {
	fmt.Println("dockit launch requires exactly 1 argument.")
	fmt.Println("See 'dockit launch --help'.")
	launch.MainHelp()
}

func (launch *Launch) MainHelp() {
	fmt.Println("\nUsage: dockit launch [OPTIONS] NAME[:TAG|@DIGEST]")
	fmt.Println("\nPull the docker image and start the container")
	fmt.Println("\nOptions:\n      --add-host list                  Add a custom host-to-IP mapping (host:ip)\n  -a, --attach list                    Attach to STDIN, STDOUT or STDERR\n      --blkio-weight uint16            Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)\n      --blkio-weight-device list       Block IO weight (relative device weight) (default [])\n      --cap-add list                   Add Linux capabilities\n      --cap-drop list                  Drop Linux capabilities\n      --cgroup-parent string           Optional parent cgroup for the container\n      --cgroupns string                Cgroup namespace to use (host|private)\n                                       'host':    Run the container in the Docker host's cgroup namespace\n                                       'private': Run the container in its own private cgroup namespace\n                                       '':        Use the cgroup namespace as configured by the\n                                                  default-cgroupns-mode option on the daemon (default)\n      --cidfile string                 Write the container ID to the file\n      --cpu-period int                 Limit CPU CFS (Completely Fair Scheduler) period\n      --cpu-quota int                  Limit CPU CFS (Completely Fair Scheduler) quota\n      --cpu-rt-period int              Limit CPU real-time period in microseconds\n      --cpu-rt-runtime int             Limit CPU real-time runtime in microseconds\n  -c, --cpu-shares int                 CPU shares (relative weight)\n      --cpus decimal                   Number of CPUs\n      --cpuset-cpus string             CPUs in which to allow execution (0-3, 0,1)\n      --cpuset-mems string             MEMs in which to allow execution (0-3, 0,1)\n  -d, --detach                         Run container in background and print container ID\n      --detach-keys string             Override the key sequence for detaching a container\n      --device list                    Add a host device to the container\n      --device-cgroup-rule list        Add a rule to the cgroup allowed devices list\n      --device-read-bps list           Limit read rate (bytes per second) from a device (default [])\n      --device-read-iops list          Limit read rate (IO per second) from a device (default [])\n      --device-write-bps list          Limit write rate (bytes per second) to a device (default [])\n      --device-write-iops list         Limit write rate (IO per second) to a device (default [])\n      --disable-content-trust          Skip image verification (default true)\n      --dns list                       Set custom DNS servers\n      --dns-option list                Set DNS options\n      --dns-search list                Set custom DNS search domains\n      --domainname string              Container NIS domain name\n      --entrypoint string              Overwrite the default ENTRYPOINT of the image\n  -e, --env list                       Set environment variables\n      --env-file list                  Read in a file of environment variables\n      --expose list                    Expose a port or a range of ports\n      --gpus gpu-request               GPU devices to add to the container ('all' to pass all GPUs)\n      --group-add list                 Add additional groups to join\n      --health-cmd string              Command to run to check health\n      --health-interval duration       Time between running the check (ms|s|m|h) (default 0s)\n      --health-retries int             Consecutive failures needed to report unhealthy\n      --health-start-period duration   Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)\n      --health-timeout duration        Maximum time to allow one check to run (ms|s|m|h) (default 0s)\n      --help                           Print usage\n  -h, --hostname string                Container host name\n      --init                           Run an init inside the container that forwards signals and reaps processes\n  -i, --interactive                    Keep STDIN open even if not attached\n      --ip string                      IPv4 address (e.g., 172.30.100.104)\n      --ip6 string                     IPv6 address (e.g., 2001:db8::33)\n      --ipc string                     IPC mode to use\n      --isolation string               Container isolation technology\n      --kernel-memory bytes            Kernel memory limit\n  -l, --label list                     Set meta data on a container\n      --label-file list                Read in a line delimited file of labels\n      --link list                      Add link to another container\n      --link-local-ip list             Container IPv4/IPv6 link-local addresses\n      --log-driver string              Logging driver for the container\n      --log-opt list                   Log driver options\n      --mac-address string             Container MAC address (e.g., 92:d0:c6:0a:29:33)\n  -m, --memory bytes                   Memory limit\n      --memory-reservation bytes       Memory soft limit\n      --memory-swap bytes              Swap limit equal to memory plus swap: '-1' to enable unlimited swap\n      --memory-swappiness int          Tune container memory swappiness (0 to 100) (default -1)\n      --mount mount                    Attach a filesystem mount to the container\n      --name string                    Assign a name to the container\n      --network network                Connect a container to a network\n      --network-alias list             Add network-scoped alias for the container\n      --no-healthcheck                 Disable any container-specified HEALTHCHECK\n      --oom-kill-disable               Disable OOM Killer\n      --oom-score-adj int              Tune host's OOM preferences (-1000 to 1000)\n      --pid string                     PID namespace to use\n      --pids-limit int                 Tune container pids limit (set -1 for unlimited)\n      --platform string                Set platform if server is multi-platform capable\n      --privileged                     Give extended privileges to this container\n  -p, --publish list                   Publish a container's port(s) to the host\n  -P, --publish-all                    Publish all exposed ports to random ports\n      --pull string                    Pull image before running (\"always\"|\"missing\"|\"never\") (default \"missing\")\n      --read-only                      Mount the container's root filesystem as read only\n      --restart string                 Restart policy to apply when a container exits (default \"no\")\n      --rm                             Automatically remove the container when it exits\n      --runtime string                 Runtime to use for this container\n      --security-opt list              Security Options\n      --shm-size bytes                 Size of /dev/shm\n      --sig-proxy                      Proxy received signals to the process (default true)\n      --stop-signal string             Signal to stop a container (default \"SIGTERM\")\n      --stop-timeout int               Timeout (in seconds) to stop a container\n      --storage-opt list               Storage driver options for the container\n      --sysctl map                     Sysctl options (default map[])\n      --tmpfs list                     Mount a tmpfs directory\n  -t, --tty                            Allocate a pseudo-TTY\n      --ulimit ulimit                  Ulimit options (default [])\n  -u, --user string                    Username or UID (format: <name|uid>[:<group|gid>])\n      --userns string                  User namespace to use\n      --uts string                     UTS namespace to use\n  -v, --volume list                    Bind mount a volume\n      --volume-driver string           Optional volume driver for the container\n      --volumes-from list              Mount volumes from the specified container(s)\n  -w, --workdir string                 Working directory inside the container")
}

func (launch *Launch) CustomExec(args []string) {
	// Pull image.
	image := docker.BuildImageModel(args[0])
	docker.Pull(image)

	// Run container.
	container := docker.BuildContainerModel(image)
	docker.Run(container, args[1:])

	// Render table.
	container.RenderTable()
}
