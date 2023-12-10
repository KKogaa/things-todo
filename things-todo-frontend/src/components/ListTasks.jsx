import { List, ListItemButton, ListItemText } from '@mui/material'

export const ListTasks = () => {

  const listData = [
    {
      name: 'Task name',
      description: 'This is the first',
      duration: '20h',
      priority: 'High',
    },
    {
      name: 'Task name',
      description: 'This is the first',
      duration: '20h',
      priority: 'High',
    },
  ]

  return (
    <List>
      {
        listData.map((task, index) => (
          <ListItemButton key={index}>
            <ListItemText primary={task.name}
              secondary={
                <div>
                  <div>
                    {task.description}
                  </div>
                  <div>
                    <strong>Duration: </strong>
                    {task.duration}
                  </div>
                  <div>
                    <strong>Priority: </strong>
                    {task.priority}
                  </div>
                </div>
              } />
          </ListItemButton>
        ))
      }
    </List>
  )
}
