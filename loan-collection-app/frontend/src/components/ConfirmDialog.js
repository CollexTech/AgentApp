import React from 'react';
import { Dialog } from '@mui/material';

const ConfirmDialog = ({ children, ...props }) => {
  return (
    <Dialog {...props}>
      {children}
    </Dialog>
  );
};

export default ConfirmDialog; 