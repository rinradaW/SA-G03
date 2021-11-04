import React, { useEffect ,useState } from "react";

import { BrowserRouter as Router, Switch, Route , Link } from "react-router-dom";

import {  createStyles,  makeStyles,  useTheme,  Theme,} from "@material-ui/core/styles";

import clsx from "clsx";

import CssBaseline from '@mui/material/CssBaseline';

import AppBar from "@material-ui/core/AppBar";

import Toolbar from "@material-ui/core/Toolbar";

import IconButton from '@mui/material/IconButton';

import MenuIcon from "@material-ui/icons/Menu";

import Avatar from '@material-ui/core/Avatar';

import Button from '@material-ui/core/Button';

import Typography from "@material-ui/core/Typography";

import Grid from "@material-ui/core/Grid";

import Tooltip from '@material-ui/core/Tooltip';

import List from "@material-ui/core/List";

import Divider from "@material-ui/core/Divider";

import Drawer from '@material-ui/core/Drawer';

import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";

import ChevronRightIcon from "@material-ui/icons/ChevronRight";

import ListItem from '@mui/material/ListItem';

import ListItemIcon from '@mui/material/ListItemIcon';

import ListItemButton from '@mui/material/ListItemButton';

import ListItemText from "@material-ui/core/ListItemText";

import Container from "@material-ui/core/Container";

import Home from "./components/Home";

import JoinActHis from "./components/JoinActHis";

import JoinActHisCreate from "./components/JoinActHisCreate";

import SignIn from "./components/SignIn";


import HomeIcon from "@material-ui/icons/Home";

import SettingsBackupRestoreIcon from '@mui/icons-material/SettingsBackupRestore';

import AccountCircleIcon from "@material-ui/icons/AccountCircle";

import { ClubCommitteesInterface } from "./models/IClubCommittee";

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: "flex",
    },

    navlink: { color: "white", textDecoration: "none" },

    tooltip: {
      backgroundColor: '#f5f5f9',
      color: 'rgba(0, 0, 0, 0.87)',
      maxWidth: 3000,
      fontSize: theme.typography.pxToRem(12),
      border: '1px solid #dadde9',
    },

    center: {
      display: "flex",
      justifyContent: "center",
      alignItems: 'center',
  
    },

    large: {
      width: theme.spacing(12),
      height: theme.spacing(12),
      margin: theme.spacing(2)
    },

    container: { margin: theme.spacing(1) },

    toolbar: {
      display: "flex",
      alignItems: "center",
      justifyContent: "flex-end",
      padding: theme.spacing(0, 1),
      // necessary for content to be below app bar
      ...theme.mixins.toolbar,
    },

    content: {
      flexGrow: 1,
      padding: theme.spacing(3),
    },

    title: {
      flexGrow: 1,
    },
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },

    appBarShift: {
      marginLeft: drawerWidth,
      width: `calc(100% - ${drawerWidth}px)`,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },

    menuButton: {
      marginRight: 36,
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
    
    a: {
      textDecoration: "none",
      color: "inherit",
    },

    
    
  })
);


export default function MiniDrawer() {

  const classes = useStyles();

  const theme = useTheme();
  
  const [open, setOpen] = React.useState(false);

  const [token, setToken] = React.useState<String>("");

  const [loginUser, setCurrentUser] = useState<ClubCommitteesInterface>();

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const menu = [
    { name: "หน้าแรก", icon: <HomeIcon />, path: "/" },
    { name: "ประวัติการเข้าร่วมกิจกรรมชมรม", icon: <AccountCircleIcon />,
     path: "/join_activity_histories" },
    
  ];

  const apiUrl = "http://localhost:8080";

  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
      
    },
  };

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }

    const loginUser = localStorage.getItem("uid");
    fetch(`${apiUrl}/club_committee/${loginUser}`, requestOptions)
    .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCurrentUser(res.data);
          
        } else {
          console.log("error loading current user");
        }
      });
  

  }, []);

  if (!token) {
    return <SignIn />;
  }

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };

 return (
  <div className={classes.root}>
   <Router>
   
        <CssBaseline />
        {token && (
          <>
            <AppBar
              position="fixed"
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
                  className={clsx(classes.menuButton, {
                    [classes.hide]: open,
                  })}
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

                  <Avatar className={classes.large} > {loginUser?.Name.substring(0, 1)} </Avatar>
                </Grid>

                <Grid className={classes.center}>
                  <Typography variant="subtitle2" gutterBottom> Login as: Committee</Typography>
                </Grid>

                <Grid className={classes.center}>
                  <Typography variant="subtitle1" gutterBottom
                  > 
                  {loginUser?.Name}
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

            <Avatar className={classes.container} > {loginUser?.Name.substring(0, 1)} </Avatar>
          </Tooltip>

        </Toolbar>
      </AppBar>
          
      <Drawer
              variant="permanent"
              className={clsx(classes.drawer, {
                [classes.drawerOpen]: open,
                [classes.drawerClose]: !open,
              })}
              classes={{
                paper: clsx({
                  [classes.drawerOpen]: open,
                  [classes.drawerClose]: !open,
                }),
              }}
            >
              <div className={classes.toolbar}>
                <IconButton onClick={handleDrawerClose}>
                  {theme.direction === "rtl" ? (
                    <ChevronRightIcon />
                  ) : (
                    <ChevronLeftIcon />
                  )}
                </IconButton>
              </div>
              <Divider />
              <List>
                {menu.map((item, index) => (
                  <Link to={item.path} key={item.name} className={classes.a}>
                    <ListItem button>
                      <ListItemIcon>{item.icon}</ListItemIcon>
                      <ListItemText primary={item.name} />
                    </ListItem>
                  </Link>
                ))}
              </List>
            </Drawer>
          </>
        )}

          <main className={classes.content}>
          <div className={classes.toolbar} />
          <div>
       
       <Switch>

         <Route exact path="/" component={Home} />
         <Route exact path="/join_activity_histories" component={JoinActHis} />
         <Route exact path="/join_activity_history/create" component={JoinActHisCreate} />

         

       </Switch>

     </div>
    
    </main>

   </Router>

   </div>

 );

}