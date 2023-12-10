import { Button } from '@mui/material';
import ListTasks from './components/ListTasks';

const App = () => {
  return (
    <div>
      <h1>Things Todo</h1>
      <Button variant="outlined">New task</Button>
      <ListTasks />
    </div>
  );
};

export default App
