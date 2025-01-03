import React from 'react';
import { 
  Drawer, 
  List, 
  ListItem, 
  ListItemText, 
  ListItemIcon,
  Box,
  Typography
} from '@mui/material';
import { 
  Home as HomeIcon, 
  ListAlt as CaseIcon, 
  People as UsersIcon,
  Business as BuildingIcon,
  GroupAdd as GroupAddIcon,
  Upload as UploadFile
} from '@mui/icons-material';
import { Link } from 'react-router-dom';

const drawerWidth = 240;

function SideNavigation({ permissions }) {
  const navigationItems = [
    {
      path: '/agency-management',
      icon: 'building',
      label: 'Agency Management'
    }
  ];

  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' },
      }}
    >
      <Box sx={{ overflow: 'auto' }}>
        {/* Navigation Section */}
        <Typography variant="h6" sx={{ p: 2 }}>
          Navigation
        </Typography>
        <List>
          <ListItem button component={Link} to="/">
            <ListItemIcon>
              <HomeIcon />
            </ListItemIcon>
            <ListItemText primary="Home" />
          </ListItem>
          <ListItem button component={Link} to="/cases">
            <ListItemIcon>
              <CaseIcon />
            </ListItemIcon>
            <ListItemText primary="My Cases" />
          </ListItem>
          <ListItem button component={Link} to="/users">
            <ListItemIcon>
              <UsersIcon />
            </ListItemIcon>
            <ListItemText primary="User Management" />
          </ListItem>
          <ListItem button component={Link} to="/agency-management">
            <ListItemIcon>
              <BuildingIcon />
            </ListItemIcon>
            <ListItemText primary="Agency Management" />
          </ListItem>
          <ListItem button component={Link} to="/agency-user-mapping">
            <ListItemIcon>
              <GroupAddIcon />
            </ListItemIcon>
            <ListItemText primary="Agency User Mapping" />
          </ListItem>
          <ListItem button component={Link} to="/case-onboarding">
            <ListItemIcon>
              <UploadFile />
            </ListItemIcon>
            <ListItemText primary="Case Onboarding" />
          </ListItem>
        </List>
      </Box>
    </Drawer>
  );
}

export default SideNavigation;