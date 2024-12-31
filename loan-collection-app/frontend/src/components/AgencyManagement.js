import React, { useState, useEffect } from 'react';
import { 
  Table, 
  TableBody, 
  TableCell, 
  TableContainer, 
  TableHead, 
  TableRow, 
  Paper, 
  Button, 
  TextField, 
  Dialog, 
  DialogTitle, 
  DialogContent, 
  DialogActions, 
  IconButton,
  Menu,
  MenuItem,
  Typography
} from '@mui/material';
import { createAgency, getAgencies, deleteAgency } from '../service/api';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import ConfirmDialog from './ConfirmDialog';

const AgencyManagement = () => {
  const [agencies, setAgencies] = useState([]);
  const [openDialog, setOpenDialog] = useState(false);
  const [errors, setErrors] = useState({
    email: '',
    phone: ''
  });
  const [newAgency, setNewAgency] = useState({
    agency_name: '',
    address: '',
    phone: '',
    email: '',
    status: 'ACTIVE'
  });
  const [anchorEl, setAnchorEl] = useState(null);
  const [selectedAgency, setSelectedAgency] = useState(null);
  const [deleteConfirmOpen, setDeleteConfirmOpen] = useState(false);

  useEffect(() => {
    fetchAgencies();
  }, []);

  const fetchAgencies = async () => {
    try {
      const response = await getAgencies();
      setAgencies(response.data);
    } catch (error) {
      console.error('Error fetching agencies:', error);
    }
  };

  const validateEmail = (email) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  };

  const validatePhone = (phone) => {
    // Allows formats: (123) 456-7890, 123-456-7890, 1234567890
    const phoneRegex = /^\(?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$/;
    return phoneRegex.test(phone);
  };

  const handleInputChange = (field, value) => {
    setNewAgency({ ...newAgency, [field]: value });
    
    // Clear error when user starts typing
    if (errors[field]) {
      setErrors({ ...errors, [field]: '' });
    }

    // Validate on change
    if (field === 'email' && value) {
      if (!validateEmail(value)) {
        setErrors(prev => ({ ...prev, email: 'Please enter a valid email address' }));
      }
    }
    if (field === 'phone' && value) {
      if (!validatePhone(value)) {
        setErrors(prev => ({ ...prev, phone: 'Please enter a valid phone number' }));
      }
    }
  };

  const handleCreateAgency = async (e) => {
    e.preventDefault();
    
    // Validate before submission
    const newErrors = {};
    if (newAgency.email && !validateEmail(newAgency.email)) {
      newErrors.email = 'Please enter a valid email address';
    }
    if (newAgency.phone && !validatePhone(newAgency.phone)) {
      newErrors.phone = 'Please enter a valid phone number';
    }

    if (Object.keys(newErrors).length > 0) {
      setErrors(newErrors);
      return;
    }

    try {
      await createAgency(newAgency);
      setOpenDialog(false);
      fetchAgencies();
      setNewAgency({ agency_name: '', address: '', phone: '', email: '', status: 'ACTIVE' });
      setErrors({ email: '', phone: '' });
    } catch (error) {
      console.error('Error creating agency:', error);
    }
  };

  const handleMenuOpen = (event, agency) => {
    setAnchorEl(event.currentTarget);
    setSelectedAgency(agency);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
    setSelectedAgency(null);
  };

  const handleDeleteAgency = async () => {
    if (!selectedAgency) return;
    
    try {
      await deleteAgency(selectedAgency.id);
      fetchAgencies(); // Refresh the list
      handleMenuClose();
    } catch (error) {
      console.error('Error deleting agency:', error);
    }
  };

  const handleDeleteConfirm = () => {
    handleDeleteAgency();
    setDeleteConfirmOpen(false);
  };

  const handleDeleteCancel = () => {
    setDeleteConfirmOpen(false);
  };

  return (
    <div className="agency-management">
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '20px' }}>
        <h2>Agency Management</h2>
        <Button 
          variant="contained" 
          color="primary" 
          onClick={() => setOpenDialog(true)}
        >
          Add New Agency
        </Button>
      </div>

      {/* Agencies Table */}
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Agency Name</TableCell>
              <TableCell>Address</TableCell>
              <TableCell>Phone</TableCell>
              <TableCell>Email</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {agencies.map((agency) => (
              <TableRow key={agency.id}>
                <TableCell>{agency.agency_name}</TableCell>
                <TableCell>{agency.address}</TableCell>
                <TableCell>{agency.phone}</TableCell>
                <TableCell>{agency.email}</TableCell>
                <TableCell>{agency.status}</TableCell>
                <TableCell>
                  <IconButton
                    size="small"
                    onClick={(e) => handleMenuOpen(e, agency)}
                  >
                    <MoreVertIcon />
                  </IconButton>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      {/* Create Agency Dialog */}
      <Dialog open={openDialog} onClose={() => setOpenDialog(false)}>
        <DialogTitle>Create New Agency</DialogTitle>
        <form onSubmit={handleCreateAgency}>
          <DialogContent>
            <TextField
              fullWidth
              label="Agency Name"
              value={newAgency.agency_name}
              onChange={(e) => handleInputChange('agency_name', e.target.value)}
              margin="normal"
              required
            />
            <TextField
              fullWidth
              label="Address"
              value={newAgency.address}
              onChange={(e) => handleInputChange('address', e.target.value)}
              margin="normal"
            />
            <TextField
              fullWidth
              label="Phone"
              value={newAgency.phone}
              onChange={(e) => handleInputChange('phone', e.target.value)}
              margin="normal"
              error={!!errors.phone}
              helperText={errors.phone}
            />
            <TextField
              fullWidth
              label="Email"
              type="email"
              value={newAgency.email}
              onChange={(e) => handleInputChange('email', e.target.value)}
              margin="normal"
              error={!!errors.email}
              helperText={errors.email}
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={() => {
              setOpenDialog(false);
              setErrors({ email: '', phone: '' });
              setNewAgency({ agency_name: '', address: '', phone: '', email: '', status: 'ACTIVE' });
            }}>
              Cancel
            </Button>
            <Button 
              type="submit" 
              variant="contained" 
              color="primary"
              disabled={!!errors.email || !!errors.phone}
            >
              Create
            </Button>
          </DialogActions>
        </form>
      </Dialog>

      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleMenuClose}
      >
        <MenuItem 
          onClick={() => setDeleteConfirmOpen(true)}
          sx={{ color: 'error.main' }}
        >
          Delete
        </MenuItem>
      </Menu>

      {/* Update the Confirmation Dialog */}
      <ConfirmDialog
        open={deleteConfirmOpen}
        onClose={handleDeleteCancel}
        maxWidth="xs"
        fullWidth
      >
        <DialogTitle sx={{ pb: 1 }}>
          Delete Agency
        </DialogTitle>
        <DialogContent sx={{ pb: 2 }}>
          <Typography>
            Are you sure you want to delete agency <strong>"{selectedAgency?.agency_name}"</strong>?
          </Typography>
          <Typography color="error" sx={{ mt: 1 }}>
            This action cannot be undone.
          </Typography>
        </DialogContent>
        <DialogActions sx={{ px: 3, pb: 2 }}>
          <Button 
            onClick={handleDeleteCancel}
            variant="outlined"
          >
            Cancel
          </Button>
          <Button 
            onClick={handleDeleteConfirm}
            color="error"
            variant="contained"
            autoFocus
          >
            Delete Agency
          </Button>
        </DialogActions>
      </ConfirmDialog>
    </div>
  );
};

export default AgencyManagement; 