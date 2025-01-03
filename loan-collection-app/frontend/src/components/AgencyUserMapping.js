import React, { useState, useEffect } from 'react';
import {
  Container,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  CircularProgress,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  IconButton,
  Grid
} from '@mui/material';
import { MoreVert, AddCircleOutline } from '@mui/icons-material';
import { getAgencies, getUnassignedUsers, getAgencyUsers, assignUserToAgency } from '../service/api';

const ROLES = ['agent', 'manager', 'supervisor'];

function AgencyUserMapping() {
  const [agencies, setAgencies] = useState([]);
  const [selectedAgency, setSelectedAgency] = useState('');
  const [unassignedUsers, setUnassignedUsers] = useState([]);
  const [agencyUsers, setAgencyUsers] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isAssignDialogOpen, setIsAssignDialogOpen] = useState(false);
  const [assignmentData, setAssignmentData] = useState({
    user_id: '',
    agency_role: 'agent'
  });

  // Load initial data
  useEffect(() => {
    loadData();
  }, []);

  // Load agency users when agency is selected
  useEffect(() => {
    if (selectedAgency) {
      loadAgencyUsers(selectedAgency);
    }
  }, [selectedAgency]);

  const loadData = async () => {
    try {
      setIsLoading(true);
      const [fetchedAgencies, fetchedUnassignedUsers] = await Promise.all([
        getAgencies(),
        getUnassignedUsers()
      ]);
      
      // Check if fetchedAgencies has the expected structure
      if (Array.isArray(fetchedAgencies?.data)) {
        setAgencies(fetchedAgencies.data);
      } else {
        console.error('Unexpected agencies data structure:', fetchedAgencies);
        setError('Error loading agencies data');
      }

      // Check if fetchedUnassignedUsers has the expected structure
      if (Array.isArray(fetchedUnassignedUsers?.data)) {
        setUnassignedUsers(fetchedUnassignedUsers.data);
      } else if (Array.isArray(fetchedUnassignedUsers)) {
        setUnassignedUsers(fetchedUnassignedUsers);
      } else {
        console.error('Unexpected unassigned users data structure:', fetchedUnassignedUsers);
        setError('Error loading unassigned users data');
      }
    } catch (err) {
      console.error('Error loading data:', err);
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const loadAgencyUsers = async (agencyId) => {
    try {
      setIsLoading(true);
      const response = await getAgencyUsers(agencyId);
      // Handle both array and wrapped response formats
      const users = Array.isArray(response) ? response : 
                   Array.isArray(response?.data) ? response.data : [];
      setAgencyUsers(users);
    } catch (err) {
      console.error('Error loading agency users:', err);
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const handleAssignUser = async () => {
    try {
      await assignUserToAgency({
        agency_id: selectedAgency,
        user_id: assignmentData.user_id,
        agency_role: assignmentData.agency_role
      });
      
      // Refresh data
      await Promise.all([
        loadAgencyUsers(selectedAgency),
        loadData()
      ]);
      
      handleAssignDialogClose();
    } catch (err) {
      console.error('Error assigning user:', err);
      setError(err.message);
    }
  };

  const handleAssignDialogOpen = () => {
    setIsAssignDialogOpen(true);
  };

  const handleAssignDialogClose = () => {
    setIsAssignDialogOpen(false);
    setAssignmentData({
      user_id: '',
      agency_role: 'agent'
    });
  };

  if (isLoading && !selectedAgency) {
    return (
      <Container sx={{ display: 'flex', justifyContent: 'center', mt: 4 }}>
        <CircularProgress />
      </Container>
    );
  }

  if (error) {
    return (
      <Container>
        <Typography color="error">Error: {error}</Typography>
      </Container>
    );
  }

  return (
    <Container>
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <Typography variant="h4" gutterBottom>
            Agency User Mapping
          </Typography>
        </Grid>

        <Grid item xs={12} md={6}>
          <FormControl fullWidth>
            <InputLabel>Select Agency</InputLabel>
            <Select
              value={selectedAgency}
              onChange={(e) => setSelectedAgency(e.target.value)}
              label="Select Agency"
            >
              {Array.isArray(agencies) && agencies.map((agency) => (
                <MenuItem key={agency.id} value={agency.id}>
                  {agency.agency_name || agency.username}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>

        <Grid item xs={12}>
          <Button
            startIcon={<AddCircleOutline />}
            onClick={handleAssignDialogOpen}
            variant="contained"
            disabled={!selectedAgency}
          >
            Assign User to Agency
          </Button>
        </Grid>

        <Grid item xs={12}>
          <TableContainer component={Paper}>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell>Email</TableCell>
                  <TableCell>Role</TableCell>
                  <TableCell>Actions</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {Array.isArray(agencyUsers) && agencyUsers.length > 0 ? (
                  agencyUsers.map((user) => (
                    <TableRow key={user.id}>
                      <TableCell>{user.username || user.name}</TableCell>
                      <TableCell>{user.email}</TableCell>
                      <TableCell>{user.role || user.agency_role}</TableCell>
                      <TableCell>
                        <IconButton>
                          <MoreVert />
                        </IconButton>
                      </TableCell>
                    </TableRow>
                  ))
                ) : (
                  <TableRow>
                    <TableCell colSpan={4} align="center">
                      No users assigned to this agency
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </TableContainer>
        </Grid>
      </Grid>

      {/* Assign User Dialog */}
      <Dialog open={isAssignDialogOpen} onClose={handleAssignDialogClose}>
        <DialogTitle>Assign User to Agency</DialogTitle>
        <DialogContent>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel>Select User</InputLabel>
            <Select
              value={assignmentData.user_id}
              onChange={(e) => setAssignmentData({ ...assignmentData, user_id: e.target.value })}
              label="Select User"
            >
              {Array.isArray(unassignedUsers) && unassignedUsers.map((user) => (
                <MenuItem key={user.id} value={user.id}>
                  {user.username || user.name} {user.email && `(${user.email})`}
                </MenuItem>
              ))}
            </Select>
          </FormControl>

          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel>Select Role</InputLabel>
            <Select
              value={assignmentData.agency_role}
              onChange={(e) => setAssignmentData({ ...assignmentData, agency_role: e.target.value })}
              label="Select Role"
            >
              {ROLES.map((role) => (
                <MenuItem key={role} value={role}>
                  {role.charAt(0).toUpperCase() + role.slice(1)}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleAssignDialogClose}>Cancel</Button>
          <Button
            onClick={handleAssignUser}
            variant="contained"
            disabled={!assignmentData.user_id || !assignmentData.agency_role}
          >
            Assign
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
}

export default AgencyUserMapping;