import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import Typography from 'material-ui/Typography';
import IconButton from 'material-ui/IconButton';
import MenuIcon from 'material-ui-icons/Menu';
import Grid from 'material-ui/Grid';
import Button from 'material-ui/Button';
import Tabs, { Tab } from 'material-ui/Tabs';

const styles = {
  card: {
    minWidth: 275,
  },
  root: {
    paddingTop: '80px',
    flexGrow: 1,
  },
  flex: {
    flex: 1,
  },
  container: {
    maringTop: 20,
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
};

function ButtonAppBar(props) {
  const { classes } = props;
  return (
    <div className={classes.root}>
      <AppBar>
        <Toolbar>
          <IconButton className={classes.menuButton} color="inherit" aria-label="Menu">
            <MenuIcon />
          </IconButton>
          <Typography variant="title" color="inherit" className={classes.flex}>
            Title
          </Typography>
          <Button color="inherit">Login</Button>
        </Toolbar>
        <Tabs>
          <Tab label="Item One" />
          <Tab label="Item Two" />
          <Tab label="Item Three" href="#basic-tabs" />
        </Tabs>
      </AppBar>
      <Grid className={classes.container} container spacing={24}>
        { [0, 1, 2, 3, 4].map(() => (
          <Grid item xs={12} sm={6}>
            <Button color="primary" className={classes.button}>
              <div>
                <div>Hello</div>
                <div>헬 로 우ㅜㄴㅁ이루루ㅏㄴㅇㅁㄹ</div>
              </div>
            </Button>
          </Grid>
        ))}
      </Grid>
    </div>
  );
}

ButtonAppBar.propTypes = {
  classes: PropTypes.shape({}).isRequired,
};

export default withStyles(styles)(ButtonAppBar);
