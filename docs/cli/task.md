# Task Management

The `task` command allows you to manage computing tasks on your provider.

## Overview

```bash
computing-provider task <subcommand> [flags]
```

## Subcommands

### List Tasks

List all tasks on your provider.

```bash
computing-provider task list [flags]
```

#### Flags

- `--fcp`: List FCP (Fog Computing Provider) tasks
- `--ecp`: List ECP (Edge Computing Provider) tasks
- `--show-completed`: Display completed jobs
- `--verbose, -v`: Show detailed information
- `--tail <number>`: Show the last N lines (default: all)

#### Examples

```bash
# List FCP tasks
computing-provider task list --fcp

# List ECP tasks with verbose output
computing-provider task list --ecp --verbose

# Show only recent FCP tasks
computing-provider task list --fcp --tail 10

# Show completed tasks
computing-provider task list --fcp --show-completed
```

#### Output Format

The task list displays the following information:

- **Job UUID**: Unique identifier for the task
- **Owner Address**: Task owner's wallet address
- **Task Type**: Type of computing task (GPU, CPU, etc.)
- **Network**: Network identifier (e.g., fil-c2)
- **Status**: Current task status
- **Verified**: Whether the task has been verified
- **Create Time**: Task creation timestamp

### Get Task Details

Get detailed information about a specific task.

```bash
computing-provider task get [job_uuid] [flags]
```

#### Arguments

- `job_uuid`: The unique identifier of the task

#### Flags

- `--fcp`: Specify FCP task
- `--ecp`: Specify ECP task

#### Examples

```bash
# Get FCP task details
computing-provider task get 1114203 --fcp

# Get ECP task details
computing-provider task get abc123 --ecp
```

#### Output Information

Task details include:

- **Basic Information**: Job UUID, owner, type, network
- **Status Information**: Current status, verification status
- **Resource Information**: CPU, memory, GPU requirements
- **Timing Information**: Creation time, start time, completion time
- **Container Information**: Container status, logs, metrics

### Delete Task

Delete a specific task from your provider.

```bash
computing-provider task delete [job_uuid] [flags]
```

#### Arguments

- `job_uuid`: The unique identifier of the task to delete

#### Flags

- `--fcp`: Specify FCP task
- `--ecp`: Specify ECP task
- `--force`: Force deletion without confirmation

#### Examples

```bash
# Delete FCP task
computing-provider task delete 1114203 --fcp

# Force delete ECP task
computing-provider task delete abc123 --ecp --force
```

## Task Statuses

### FCP Task Statuses

- **Pending**: Task is waiting to be scheduled
- **Running**: Task is currently executing
- **Completed**: Task has finished successfully
- **Failed**: Task has failed
- **Cancelled**: Task was cancelled
- **Timeout**: Task exceeded time limit

### ECP Task Statuses

- **Submitted**: Task has been submitted to the network
- **Processing**: Task is being processed
- **Completed**: Task has been completed successfully
- **Failed**: Task has failed
- **Verified**: Task has been verified on-chain

## Task Types

### FCP Task Types

- **AI Training**: Machine learning model training
- **AI Inference**: Model inference and prediction
- **Data Processing**: Large-scale data processing
- **Custom**: Custom container workloads

### ECP Task Types

- **fil-c2**: Filecoin C2 ZK-SNARK proof generation
- **zk-proof**: General zero-knowledge proof generation

## Resource Requirements

### CPU Tasks
- **CPU Cores**: Number of CPU cores required
- **Memory**: RAM requirements in GB
- **Storage**: Disk space requirements

### GPU Tasks
- **GPU Count**: Number of GPUs required
- **GPU Memory**: GPU memory requirements
- **CUDA Version**: Required CUDA version

## Monitoring Tasks

### Real-time Monitoring

```bash
# Monitor FCP tasks in real-time
computing-provider task list --fcp --tail 5

# Watch for new tasks
watch -n 5 'computing-provider task list --fcp --tail 1'
```

### Task Logs

```bash
# Get task logs (if supported)
computing-provider task logs [job_uuid] --fcp
```

## Task Management Best Practices

### Resource Allocation

- Monitor resource usage to avoid overloading
- Set appropriate resource limits for tasks
- Use resource quotas to manage capacity

### Task Scheduling

- Prioritize tasks based on importance
- Monitor task queue and adjust capacity
- Handle task failures gracefully

### Performance Optimization

- Monitor task completion times
- Optimize resource allocation
- Track task success rates

## Troubleshooting

### Common Issues

1. **Task Stuck in Pending**
   - Check resource availability
   - Verify Kubernetes cluster health
   - Check network connectivity

2. **Task Failures**
   - Review task logs for errors
   - Check resource constraints
   - Verify container image availability

3. **Resource Exhaustion**
   - Monitor resource usage
   - Adjust resource limits
   - Scale up infrastructure if needed

### Debug Commands

```bash
# Check provider status
computing-provider state

# Check resource usage
computing-provider info

# Check network connectivity
computing-provider network status
```

## Related Commands

- [`ubi-task`](ubi-task.md) - Manage UBI tasks
- [`info`](info.md) - Provider information
- [`state`](state.md) - Provider state
- [`network`](network.md) - Network configuration 