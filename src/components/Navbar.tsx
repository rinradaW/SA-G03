import React, { useEffect } from "react";

import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import clsx from 'clsx';

import { styled, makeStyles, useTheme } from "@material-ui/core/styles";

import Grid from "@material-ui/core/Grid";

import Box from '@mui/material/Box';

import AppBar from "@material-ui/core/AppBar";

import Toolbar from "@material-ui/core/Toolbar";

import ListItemIcon from '@mui/material/ListItemIcon';

import CssBaseline from '@mui/material/CssBaseline';

import IconButton from '@mui/material/IconButton';

import List from "@material-ui/core/List";
import Divider from "@material-ui/core/Divider";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ChevronRightIcon from "@material-ui/icons/ChevronRight";
import ListItem from '@mui/material/ListItem';

import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from "@material-ui/core/ListItemText";

import Typography from "@material-ui/core/Typography";

import MenuIcon from "@material-ui/icons/Menu";

import Drawer from '@material-ui/core/Drawer';

import Avatar from '@material-ui/core/Avatar';

import Button from '@material-ui/core/Button';

import Tooltip from '@material-ui/core/Tooltip';

import Container from "@material-ui/core/Container";

import SwipeableDrawer from '@mui/material/SwipeableDrawer';

import HomeIcon from "@material-ui/icons/Home";

import DraftsIcon from '@mui/icons-material/Drafts';

import SignIn from "./SignIn";


const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({

  root: { flexGrow: 1 },

  container: { margin: theme.spacing(1) },

  menuButton: { marginRight: theme.spacing(2) },

  boxprofile: { marginLeft: theme.spacing(2.5) },

  large: {
    width: theme.spacing(12),
    height: theme.spacing(12),
    margin: theme.spacing(2)
  },

  title: { flexGrow: 1 },

  navlink: { color: "white", textDecoration: "none" },

  tooltip: {
    backgroundColor: '#f5f5f9',
    color: 'rgba(0, 0, 0, 0.87)',
    maxWidth: 3000,
    fontSize: theme.typography.pxToRem(12),
    border: '1px solid #dadde9',
  },

  hide: {
    display: "none",
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
    whiteSpace: "nowrap",
  },
  drawerOpen: {
    width: drawerWidth,
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerClose: {
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    overflowX: "hidden",
    width: theme.spacing(7) + 1,
    [theme.breakpoints.up("sm")]: {
      width: theme.spacing(9) + 1,
    },
  },

  appBarShift: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },

  appBar: {
    transition: theme.transitions.create(['margin', 'width'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },

  content: {
    flexGrow: 1,

    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    marginLeft: -drawerWidth,
  },
  contentShift: {
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
    marginLeft: 0,
  },

  center: {
    display: "flex",
    justifyContent: "center",
    alignItems: 'center',

  },

  sx: {
    width: drawerWidth,
    flexShrink: 0,
    '& .MuiDrawer-paper': {
      width: drawerWidth,
    }
  },

  toolbar: {
    display: "flex",
    alignItems: "center",
    justifyContent: "flex-end",
    padding: theme.spacing(0, 1),
    // necessary for content to be below app bar
    ...theme.mixins.toolbar,
  },

}));


function Navbar() {

  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);
  const [token, setToken] = React.useState<String>("");
  
  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }
  }, []);

  if (!token) {
    return <SignIn />;
  }

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };


  const DrawerHeader = styled('div')(({ theme }) => ({
    display: 'flex',
    alignItems: 'center',
    padding: theme.spacing(0, 1),
    // necessary for content to be below app bar
    ...theme.mixins.toolbar,
    justifyContent: 'flex-end',
  }));

  
  return (
    <div className={classes.root}>
      <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <AppBar position="fixed"
          className={clsx(classes.appBar, {
            [classes.appBarShift]: open,
          })}
        
      >
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            edge="start"
            sx={{ mr: 2, ...(open && { display: 'none' }) }}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h5" className={classes.title}>
            <Link className={classes.navlink} to="/">
              SUT CLUB
            </Link>
          </Typography>
          
          <Typography variant="body1" >  </Typography>


          <Tooltip arrow
            className={classes.tooltip} interactive

            title={
              <Container>

                <Grid className={classes.center}>

                  <Avatar className={classes.large} > R </Avatar>
                </Grid>

                <Grid className={classes.center}>
                  <Typography variant="subtitle2" gutterBottom> Login as: Committee</Typography>
                </Grid>

                <Grid className={classes.center}>
                  <Typography variant="subtitle1" gutterBottom > Rinrada

                  </Typography>
                </Grid>


                <Grid className={classes.center}>
                  <Button className={classes.container}
                    variant="contained"
                    color="default"
                    size="medium"
                    onClick={signout}>
                    
                    Logout</Button>
                </Grid>

              </Container>

            }
          >

            <Avatar className={classes.container} > R </Avatar>
          </Tooltip>

        </Toolbar>
      </AppBar>

      <Drawer
        className={classes.sx}
        variant="persistent"
        anchor="left"
        open={open}
      >
        <DrawerHeader>
          <IconButton onClick={handleDrawerClose}>
            {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
          </IconButton>
        </DrawerHeader>
        <Divider />
        <List>
          <ListItem disablePadding>
          <Link className={classes.navlink} to="/"></Link>
            <ListItemButton>
              <ListItemIcon>
                <HomeIcon />
                
              </ListItemIcon>
              <ListItemText primary="Home" />
              
            </ListItemButton>
          </ListItem>


          <ListItem disablePadding>
          <Link className={classes.navlink} to="/join_activity_histories"></Link>
            <ListItemButton>
              
            <ListItemIcon>
                <HomeIcon />
                
              </ListItemIcon>
                
              
              <ListItemText primary="Activity History" />
              
            </ListItemButton>
          </ListItem>
        </List>
        </Drawer>
        <main className={classes.content}>
          <div className={classes.toolbar} />
          
        </main>
        </Box>

    </div>
    
  );
}

export default Navbar;