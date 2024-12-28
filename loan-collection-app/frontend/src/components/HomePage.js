import React from 'react';
import { 
  Container, 
  Typography, 
  Paper 
} from '@mui/material';

function HomePage() {
  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Paper elevation={3} sx={{ p: 4, textAlign: 'center' }}>
        <Typography variant="h4" gutterBottom>
          Welcome to Loan Collection App
        </Typography>
        <Typography variant="body1" paragraph>
          This is your central dashboard for managing loan collection cases.
        </Typography>
      </Paper>
    </Container>
  );
}

export default HomePage;