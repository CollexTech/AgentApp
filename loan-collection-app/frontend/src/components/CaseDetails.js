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
  const [caseInfo, setCaseInfo] = useState({});
  const [trails, setTrails] = useState([]);
  const [contacted, setContacted] = useState(false);
  const [paymentDate, setPaymentDate] = useState("");
  const [remarks, setRemarks] = useState("");
  const [paymentLink, setPaymentLink] = useState("");

  useEffect(() => {
    async function fetchData() {
      try {
        const info = await getCaseDetails(caseId);
        setCaseInfo(info);
        const tr = await getTrails(caseId);
        setTrails(tr);
      } catch (err) {
        console.error(err);
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

  return (
    <Box p={4}>
      <Typography variant="h4" gutterBottom>
        Case #{caseId} Details
      </Typography>
      <Paper style={{ padding: 16, marginBottom: 16 }}>
        <Typography variant="h6">User: {caseInfo.user_name}</Typography>
        <Typography>Loan ID: {caseInfo.loan_id}</Typography>
        <Typography>Loan Amount: ₹ {caseInfo.loan_amount}</Typography>
        <Typography>Monthly EMI: ₹ {caseInfo.emi_monthly}</Typography>
        <Typography>Days Past Due: {caseInfo.days_past_due}</Typography>
        <Typography>Address: {caseInfo.customer_addr}</Typography>
        <Typography>Phone: {caseInfo.customer_phone}</Typography>
        <Button variant="outlined" sx={{ mt: 2 }} onClick={openMap}>
          View on Map
        </Button>
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
        {trails.map((t) => (
          <Paper key={t.trail_id} style={{ padding: 16, marginBottom: 8 }}>
            <Typography>Trail ID: {t.trail_id}</Typography>
            <Typography>Contacted: {t.contacted ? "Yes" : "No"}</Typography>
            <Typography>Payment Date: {t.payment_date || "N/A"}</Typography>
            <Typography>Remarks: {t.remarks}</Typography>
          </Paper>
        ))}
      </Box>
    </Box>
  );
}

export default CaseDetails;
