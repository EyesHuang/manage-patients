import React, { useState, useEffect } from 'react';
import {
  List,
  ListItem,
  ListItemText,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
} from '@mui/material';
import axios from 'axios';

function App() {
  const [patients, setPatients] = useState([]);
  const [open, setOpen] = useState(false);
  const [selectedPatient, setSelectedPatient] = useState(null);
  const [orderText, setOrderText] = useState('');

  useEffect(() => {
    axios.get('http://localhost:8080/patients')
      .then((response) => {
        setPatients(response.data);
      })
      .catch((error) => console.error('There was an error!', error));
  }, []);

  const handleClickOpen = (patient) => {
    setSelectedPatient(patient);
    setOpen(true);
    setOrderText(patient.order.message);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleUpdateOrder = () => {
    const formData = new FormData();
    formData.append('message', orderText);

    axios.put(`http://localhost:8080/orders/${selectedPatient.order.id}`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })
      .then((response) => {
        const updatedOrder = response.data;

        setPatients(patients.map((patient) => {
          if (patient.id === selectedPatient.id) {
            return { ...patient, order: updatedOrder };
          }
          return patient;
        }));
        handleClose();
      })
      .catch((error) => console.error('There was an error!', error));
  };

  return (
    <div>
      <List>
        {patients.map((patient) => (
          <ListItem button key={patient.id} onClick={() => handleClickOpen(patient)}>
            <ListItemText primary={patient.name} />
          </ListItem>
        ))}
      </List>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>
          {`${selectedPatient?.name}'s Order`}
        </DialogTitle>
        <DialogContent>
          <TextField
            margin="dense"
            id="order"
            label="Order Message"
            type="text"
            fullWidth
            variant="standard"
            value={orderText}
            onChange={(e) => setOrderText(e.target.value)}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleUpdateOrder}>Update Order</Button>
          <Button onClick={handleClose}>Close</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

export default App;
