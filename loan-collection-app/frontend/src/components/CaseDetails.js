import React, { useEffect, useMemo, useState } from "react";
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
  Card,
} from "@mui/material";

const caseFields = [
  { label: "Case Status", key: "case_status", default: "N/A" },
  { label: "Agent Name", key: "agent_name", default: "N/A" },
  { label: "Loan ID", key: "loan_id", default: "N/A" },
  { label: "Loan Amount", key: "loan_amount", default: 0, prefix: "₹ " },
  { label: "Monthly EMI", key: "emi_monthly", default: 0, prefix: "₹ " },
  { label: "Days Past Due", key: "days_past_due", default: 0 },
  { label: "DPD", key: "dpd", default: 0 },
  { label: "DPD Bucket", key: "dpd_bucket", default: "N/A" },
  { 
    label: "EMI Date", 
    key: "emi_date", 
    default: "N/A", 
    format: (value) => value ? new Date(value).toLocaleDateString() : "N/A" 
  },
  { label: "Loan Description", key: "loan_description", default: "N/A" },
  { label: "EMIs Paid", key: "emis_paid_till_date", default: 0 },
  { label: "EMIs Pending", key: "emis_pending", default: 0 },
  { label: "Bounce Charges", key: "bounce_charges", default: 0, prefix: "₹ " },
  { label: "NACH Status", key: "nach_presentation_status", default: "N/A" },
  { 
    label: "Insurance Active", 
    key: "insurance_active", 
    default: "No", 
    format: (value) => value ? "Yes" : "No" 
  },
  { 
    label: "Disbursal Date", 
    key: "disbursal_date", 
    default: "N/A", 
    format: (value) => value ? new Date(value).toLocaleDateString() : "N/A" 
  }
];

const trailFields = [
  { label: "Trail ID", key: "trail_id", default: "N/A" },
  { 
    label: "Contacted", 
    key: "contacted", 
    default: "No", 
    format: (value) => value ? "Yes" : "No" 
  },
  { label: "Payment Date", key: "payment_date", default: "N/A" },
  { label: "Remarks", key: "remarks", default: "N/A" }
];

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
        if (trailsResponse?.data) {
          setTrails(trailsResponse?.data);
        }
        
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
      if (paymentDate && contacted.remarks) {
        await postTrail(caseId, newTrail);
        // Refresh trails
        const tr = await getTrails(caseId);
        setTrails(tr);
        // Clear form
        setContacted(false);
        setPaymentDate("");
        setRemarks("");
      }
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
        <div style={{display: 'flex', flexWrap: 'wrap', gap: 8}}>
        {caseFields.map(({ label, key, default: defaultValue, prefix = "", format }) => {
          const value = caseInfo[key];
          const displayValue = format ? format(value) : value ?? defaultValue;
          return (
            <Paper style={{ padding: '8px', cursor: 'pointer' }} key={key} >
              <Typography key={key} variant='body1'>
              {label}: {prefix}{displayValue}
            </Typography>
            </Paper>
          );
        })}
        </div>
        
        <Box mt={4}>
          <Typography variant="h6" style={{ marginBottom: 12 }}>Customer Contact Details</Typography>
          <Paper style={{ padding: "8px", marginBottom: '8px' }}><Typography>Address: {caseInfo.customer_addr || 'N/A'}</Typography></Paper>
          <Paper style={{ padding: "8px", marginBottom: '8px' }}><Typography>Phone: {caseInfo.customer_phone || 'N/A'}</Typography></Paper>
          {caseInfo.customer_addr && (
            <Button variant="outlined" onClick={openMap} sx={{ mt: 1 }}>
              View on Map
            </Button>
          )}
        </Box>
      </Paper>

      <Card style={{ padding: 16, marginBottom: 16 }}>
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
      </Card>

      <Card style={{ padding: 16, marginBottom: 16 }}>
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
      </Card>

      <Card style={{ padding: 16, marginBottom: 16 }}>
        <Typography variant="h5" gutterBottom>
          Trail History
        </Typography>
        {trails && trails.length > 0 ? (
          trails.map((trail) => (
            <Paper key={trail.trail_id} style={{ padding: 16, marginBottom: 8 }}>
              <div style={{display: 'flex', flexWrap: 'wrap', gap: 8}}>
                {trailFields.map(({ label, key, default: defaultValue, format }) => {
                  const value = trail[key];
                  const displayValue = format ? format(value) : value ?? defaultValue;
                  return (
                    <Paper style={{ padding: "8px", marginBottom: '8px' }}>
                      <Typography key={key}>
                        {label}: {displayValue}
                      </Typography>
                    </Paper>
                  );
                })}
              </div>
            </Paper>
          ))
        ) : (
          <Typography color="textSecondary">No trails recorded yet.</Typography>
        )}
      </Card>
    </Box>
  );
}

export default CaseDetails;
