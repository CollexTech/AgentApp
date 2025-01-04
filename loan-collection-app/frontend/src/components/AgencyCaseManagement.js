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
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Grid,
  Box,
  Alert,
  Checkbox
} from '@mui/material';
import { AssignmentInd } from '@mui/icons-material';
import { getAgencyCases, getMyAgencyUsers, assignCaseToUser } from '../service/api';

function AgencyCaseManagement() {
  const [cases, setCases] = useState([]);
  const [agencyUsers, setAgencyUsers] = useState([]);
  const [selectedCases, setSelectedCases] = useState([]);
  const [selectedUser, setSelectedUser] = useState('');
  const [isAssignDialogOpen, setIsAssignDialogOpen] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setIsLoading(true);
      const [fetchedCases, fetchedUsers] = await Promise.all([
        getAgencyCases(),
        getMyAgencyUsers()
      ]);
      
      setCases(fetchedCases.data || []);
      setAgencyUsers(fetchedUsers.data || []);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const handleAssignCases = async () => {
    try {
      setIsLoading(true);
      await Promise.all(
        selectedCases.map(caseId =>
          assignCaseToUser({
            case_id: caseId,
            user_id: selectedUser
          })
        )
      );
      setSuccess('Cases assigned successfully');
      loadData();
      setIsAssignDialogOpen(false);
      setSelectedCases([]);
      setSelectedUser('');
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const handleCaseSelection = (caseId) => {
    setSelectedCases(prev => {
      if (prev.includes(caseId)) {
        return prev.filter(id => id !== caseId);
      }
      return [...prev, caseId];
    });
  };

  return (
    <Container>
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <Typography variant="h4" gutterBottom>
            Agency Case Management
          </Typography>
        </Grid>

        {error && (
          <Grid item xs={12}>
            <Alert severity="error" onClose={() => setError(null)}>
              {error}
            </Alert>
          </Grid>
        )}

        {success && (
          <Grid item xs={12}>
            <Alert severity="success" onClose={() => setSuccess(null)}>
              {success}
            </Alert>
          </Grid>
        )}

        <Grid item xs={12}>
          <Box sx={{ mb: 2 }}>
            <Button
              variant="contained"
              startIcon={<AssignmentInd />}
              onClick={() => setIsAssignDialogOpen(true)}
              disabled={selectedCases.length === 0}
            >
              Assign to User ({selectedCases.length})
            </Button>
          </Box>
        </Grid>

        <Grid item xs={12}>
          <TableContainer component={Paper}>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell padding="checkbox">
                    <Checkbox
                      onChange={(e) => {
                        if (e.target.checked) {
                          setSelectedCases(cases.map(c => c.id));
                        } else {
                          setSelectedCases([]);
                        }
                      }}
                      checked={cases.length > 0 && selectedCases.length === cases.length}
                      indeterminate={selectedCases.length > 0 && selectedCases.length < cases.length}
                    />
                  </TableCell>
                  <TableCell>Loan ID</TableCell>
                  <TableCell>Customer ID</TableCell>
                  <TableCell>EMI Amount</TableCell>
                  <TableCell>Status</TableCell>
                  <TableCell>Assigned To</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {cases.map((case_) => (
                  <TableRow key={case_.id}>
                    <TableCell padding="checkbox">
                      <Checkbox
                        checked={selectedCases.includes(case_.id)}
                        onChange={() => handleCaseSelection(case_.id)}
                      />
                    </TableCell>
                    <TableCell>{case_.loan_id}</TableCell>
                    <TableCell>{case_.external_customer_id}</TableCell>
                    <TableCell>{case_.emi_amount}</TableCell>
                    <TableCell>{case_.case_status}</TableCell>
                    <TableCell>{case_.assigned_user?.username || 'Unassigned'}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Grid>
      </Grid>

      {/* Assign Cases Dialog */}
      <Dialog open={isAssignDialogOpen} onClose={() => setIsAssignDialogOpen(false)}>
        <DialogTitle>Assign Cases to User</DialogTitle>
        <DialogContent>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel>Select User</InputLabel>
            <Select
              value={selectedUser}
              onChange={(e) => setSelectedUser(e.target.value)}
              label="Select User"
            >
              {agencyUsers.map((user) => (
                <MenuItem key={user.id} value={user.id}>
                  {user.username}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setIsAssignDialogOpen(false)}>Cancel</Button>
          <Button
            onClick={handleAssignCases}
            variant="contained"
            disabled={!selectedUser}
          >
            Assign Cases
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
}

export default AgencyCaseManagement; 