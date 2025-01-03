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
  IconButton,
  Box,
  Alert
} from '@mui/material';
import { UploadFile, AssignmentInd, MoreVert } from '@mui/icons-material';
import { 
  uploadCases, 
  getUnassignedCases, 
  getAgencies, 
  assignCasesToAgency 
} from '../service/api';

function CaseOnboarding() {
  const [cases, setCases] = useState([]);
  const [agencies, setAgencies] = useState([]);
  const [selectedAgency, setSelectedAgency] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);
  const [isAssignDialogOpen, setIsAssignDialogOpen] = useState(false);
  const [selectedCases, setSelectedCases] = useState([]);
  const [uploadDialogOpen, setUploadDialogOpen] = useState(false);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setIsLoading(true);
      const [fetchedCases, fetchedAgencies] = await Promise.all([
        getUnassignedCases(),
        getAgencies()
      ]);
      
      setCases(fetchedCases.data || []);
      setAgencies(fetchedAgencies.data || []);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const handleFileUpload = async (event) => {
    const file = event.target.files[0];
    if (!file) return;

    const formData = new FormData();
    formData.append('file', file);

    try {
      setIsLoading(true);
      await uploadCases(formData);
      setSuccess('Cases uploaded successfully');
      loadData(); // Refresh the cases list
      setUploadDialogOpen(false);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  const handleAssignCases = async () => {
    try {
      setIsLoading(true);
      await assignCasesToAgency({
        agency_id: selectedAgency,
        case_ids: selectedCases
      });
      setSuccess('Cases assigned successfully');
      loadData(); // Refresh the list
      setIsAssignDialogOpen(false);
      setSelectedCases([]);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Container>
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <Typography variant="h4" gutterBottom>
            Case Onboarding
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
              startIcon={<UploadFile />}
              onClick={() => setUploadDialogOpen(true)}
              sx={{ mr: 2 }}
            >
              Upload Cases
            </Button>
            <Button
              variant="contained"
              startIcon={<AssignmentInd />}
              onClick={() => setIsAssignDialogOpen(true)}
              disabled={!cases.length}
            >
              Assign to Agency
            </Button>
          </Box>
        </Grid>

        <Grid item xs={12}>
          <TableContainer component={Paper}>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Loan ID</TableCell>
                  <TableCell>Customer ID</TableCell>
                  <TableCell>EMI Amount</TableCell>
                  <TableCell>Status</TableCell>
                  <TableCell>Actions</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {cases.map((case_) => (
                  <TableRow key={case_.id}>
                    <TableCell>{case_.loan_id}</TableCell>
                    <TableCell>{case_.external_customer_id}</TableCell>
                    <TableCell>{case_.emi_amount}</TableCell>
                    <TableCell>{case_.case_status}</TableCell>
                    <TableCell>
                      <IconButton>
                        <MoreVert />
                      </IconButton>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Grid>
      </Grid>

      {/* Upload Dialog */}
      <Dialog open={uploadDialogOpen} onClose={() => setUploadDialogOpen(false)}>
        <DialogTitle>Upload Cases</DialogTitle>
        <DialogContent>
          <Typography variant="body2" color="textSecondary" sx={{ mb: 2 }}>
            Please upload a CSV file with the following columns:
            loan_id, external_customer_id, emi_amount, principal_outstanding, etc.
          </Typography>
          <Button
            variant="contained"
            component="label"
          >
            Select File
            <input
              type="file"
              hidden
              accept=".csv"
              onChange={handleFileUpload}
            />
          </Button>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setUploadDialogOpen(false)}>Cancel</Button>
        </DialogActions>
      </Dialog>

      {/* Assign Cases Dialog */}
      <Dialog open={isAssignDialogOpen} onClose={() => setIsAssignDialogOpen(false)}>
        <DialogTitle>Assign Cases to Agency</DialogTitle>
        <DialogContent>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel>Select Agency</InputLabel>
            <Select
              value={selectedAgency}
              onChange={(e) => setSelectedAgency(e.target.value)}
              label="Select Agency"
            >
              {agencies.map((agency) => (
                <MenuItem key={agency.id} value={agency.id}>
                  {agency.agency_name}
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
            disabled={!selectedAgency}
          >
            Assign Cases
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
}

export default CaseOnboarding; 