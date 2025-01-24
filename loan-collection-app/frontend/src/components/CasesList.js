import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getCases } from "../service/api";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Paper,
  Typography,
  Box,
  Button,
} from "@mui/material";

function CasesList() {
  const [cases, setCases] = useState([]);
  const [totalEarnings, setTotalEarnings] = useState(0);
  const navigate = useNavigate();

  useEffect(() => {
    async function fetchCases() {
      try {
        const response = await getCases();
        setCases(response.data || []);
        setTotalEarnings(response.total_earnings || 0);
      } catch (err) {
        console.error(err);
        setCases([]);
      }
    }
    fetchCases();
  }, []);

  return (
    <Box p={4}>
      <Typography variant="h4" gutterBottom>
        Assigned Cases
      </Typography>
      <Typography variant="h6" gutterBottom>
        Total Earnings Today: ₹ {totalEarnings}
      </Typography>
      <Paper>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Case ID</TableCell>
              <TableCell>Loan ID</TableCell>
              <TableCell>Customer ID</TableCell>
              <TableCell>EMI Amount</TableCell>
              <TableCell>DPD</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {cases.map((c) => (
              <TableRow key={c.id}>
                <TableCell>{c.id}</TableCell>
                <TableCell>{c.loan_id}</TableCell>
                <TableCell>{c.external_customer_id}</TableCell>
                <TableCell>₹ {c.emi_amount}</TableCell>
                <TableCell>{c.dpd}</TableCell>
                <TableCell>{c.case_status}</TableCell>
                <TableCell>
                  <Button
                    variant="outlined"
                    onClick={() => {
                      if (c.id) {
                        navigate(`/cases/${c.id}`);
                      } else {
                        console.error('Case ID is undefined');
                      }
                    }}
                  >
                    View
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Paper>
    </Box>
  );
}

export default CasesList;
