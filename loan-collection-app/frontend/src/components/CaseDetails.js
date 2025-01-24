import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getCaseDetails, getTrails, postTrail, getPaymentLink } from "../service/api";
import {
  Typography,
  Box,
  Paper,
  Button,
  TextField,
  Checkbox,
  FormControlLabel,
} from "@mui/material";

function CaseDetails() {
  const { caseId } = useParams();
  const [caseInfo, setCaseInfo] = useState(null);
  const [trails, setTrails] = useState([]);
  const [contacted, setContacted] = useState(false);
  const [paymentDate, setPaymentDate] = useState("");
  const [remarks, setRemarks] = useState("");
  const [paymentLink, setPaymentLink] = useState("");

  useEffect(() => {
    async function fetchData() {
      if (!caseId) return;

      try {
        const response = await getCaseDetails(caseId);
        setCaseInfo(response.data);
        const trailsResponse = await getTrails(caseId);
        setTrails(trailsResponse?.data || []);
      } catch (err) {
        setCaseInfo({});
        setTrails([]);
      }
    }
    fetchData();
  }, [caseId]);

  const handleAddTrail = async () => {
    try {
      const newTrail = {
        contacted,
        payment_date: paymentDate,
        remarks,
      };
      await postTrail(caseId, newTrail);
      // Refresh trails
      const tr = await getTrails(caseId);
      setTrails(tr);
      // Clear form
      setContacted(false);
      setPaymentDate("");
      setRemarks("");
    } catch (err) {
      console.error(err);
    }
  };

  const handleGetPaymentLink = async () => {
    try {
      const linkResp = await getPaymentLink(caseId);
      setPaymentLink(linkResp.payment_link);
    } catch (err) {
      console.error(err);
    }
  };

  const openMap = () => {
    // Construct google maps search url
    const address = encodeURIComponent(caseInfo.customer_addr || "");
    window.open(`https://www.google.com/maps/search/?api=1&query=${address}`, "_blank");
  };

  if (!caseId) {
    return (
      <Box p={4}>
        <Typography variant="h4" color="error">
          Error: Case ID not found
        </Typography>
      </Box>
    );
  }

  if (!caseInfo) {
    return (
      <Box p={4}>
        <Typography>Loading case details...</Typography>
      </Box>
    );
  }

  return (
    <Box p={4}>
      <Typography variant="h4" gutterBottom>
        Case #{caseInfo.case_id} Details
      </Typography>
      <Paper style={{ padding: 16, marginBottom: 16 }}>
        <Typography variant="h6">Case Status: {caseInfo.case_status || 'N/A'}</Typography>
        <Typography>Agent Name: {caseInfo.agent_name || 'N/A'}</Typography>
        <Typography>Loan ID: {caseInfo.loan_id || 'N/A'}</Typography>
        <Typography>Loan Amount: ₹ {caseInfo.loan_amount || 0}</Typography>
        <Typography>Monthly EMI: ₹ {caseInfo.emi_monthly || 0}</Typography>
        <Typography>Days Past Due: {caseInfo.days_past_due || 0}</Typography>
        <Typography>DPD: {caseInfo.dpd || 0}</Typography>
        <Typography>DPD Bucket: {caseInfo.dpd_bucket || 'N/A'}</Typography>
        <Typography>EMI Date: {caseInfo.emi_date ? new Date(caseInfo.emi_date).toLocaleDateString() : 'N/A'}</Typography>
        <Typography>Loan Description: {caseInfo.loan_description || 'N/A'}</Typography>
        <Typography>EMIs Paid: {caseInfo.emis_paid_till_date || 0}</Typography>
        <Typography>EMIs Pending: {caseInfo.emis_pending || 0}</Typography>
        <Typography>Bounce Charges: ₹ {caseInfo.bounce_charges || 0}</Typography>
        <Typography>NACH Status: {caseInfo.nach_presentation_status || 'N/A'}</Typography>
        <Typography>Insurance Active: {caseInfo.insurance_active ? "Yes" : "No"}</Typography>
        <Typography>Disbursal Date: {caseInfo.disbursal_date ? new Date(caseInfo.disbursal_date).toLocaleDateString() : 'N/A'}</Typography>
        
        <Box mt={2}>
          <Typography variant="h6">Customer Contact Details</Typography>
          <Typography>Address: {caseInfo.customer_addr || 'N/A'}</Typography>
          <Typography>Phone: {caseInfo.customer_phone || 'N/A'}</Typography>
          {caseInfo.customer_addr && (
            <Button variant="outlined" onClick={openMap} sx={{ mt: 1 }}>
              View on Map
            </Button>
          )}
        </Box>
      </Paper>

      <Box my={2}>
        <Typography variant="h5" gutterBottom>
          Log a Trail
        </Typography>
        <FormControlLabel
          control={<Checkbox checked={contacted} onChange={(e) => setContacted(e.target.checked)} />}
          label="Contacted User"
        />
        <TextField
          label="Promised Payment Date"
          type="date"
          value={paymentDate}
          onChange={(e) => setPaymentDate(e.target.value)}
          InputLabelProps={{ shrink: true }}
          style={{ display: "block", marginBottom: "16px" }}
        />
        <TextField
          label="Remarks"
          multiline
          rows={3}
          value={remarks}
          onChange={(e) => setRemarks(e.target.value)}
          fullWidth
          style={{ marginBottom: "16px" }}
        />
        <Button variant="contained" onClick={handleAddTrail}>
          Submit Trail
        </Button>
      </Box>

      <Box my={2}>
        <Typography variant="h5" gutterBottom>
          Share Payment Link
        </Typography>
        <Button variant="contained" onClick={handleGetPaymentLink}>
          Generate Link
        </Button>
        {paymentLink && (
          <Box mt={2}>
            <Typography>Payment Link: <a href={paymentLink} target="_blank" rel="noreferrer">{paymentLink}</a></Typography>
          </Box>
        )}
      </Box>

      <Box my={2}>
        <Typography variant="h5" gutterBottom>
          Trail History
        </Typography>
        {trails && trails.length > 0 ? (
          trails.map((t) => (
            <Paper key={t.trail_id} style={{ padding: 16, marginBottom: 8 }}>
              <Typography>Trail ID: {t.trail_id}</Typography>
              <Typography>Contacted: {t.contacted ? "Yes" : "No"}</Typography>
              <Typography>Payment Date: {t.payment_date || "N/A"}</Typography>
              <Typography>Remarks: {t.remarks}</Typography>
            </Paper>
          ))
        ) : (
          <Typography color="textSecondary">No trails recorded yet.</Typography>
        )}
      </Box>
    </Box>
  );
}

export default CaseDetails;
