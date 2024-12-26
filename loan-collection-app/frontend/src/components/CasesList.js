import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getCases } from "../services/api";
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
        setCases(response.cases);
        setTotalEarnings(response.total_earnings);
      } catch (err) {
        console.error(err);
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
              <TableCell>User Name</TableCell>
              <TableCell>Loan ID</TableCell>
              <TableCell>Loan Amount</TableCell>
              <TableCell>Days Past Due</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {cases.map((c) => (
              <TableRow key={c.case_id}>
                <TableCell>{c.case_id}</TableCell>
                <TableCell>{c.user_name}</TableCell>
                <TableCell>{c.loan_id}</TableCell>
                <TableCell>{c.loan_amount}</TableCell>
                <TableCell>{c.days_past_due}</TableCell>
                <TableCell>
                  <Button
                    variant="outlined"
                    onClick={() => navigate(`/cases/${c.case_id}`)}
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
