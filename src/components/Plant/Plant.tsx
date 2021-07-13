/* eslint-disable react/jsx-no-undef */
import {
  Avatar,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  Collapse,
  Divider,
  Grid,
  IconButton,
  Typography,
} from "@material-ui/core";
import { useState } from "react";
import MoreVertIcon from "@material-ui/icons/MoreVert";
import FavoriteIcon from "@material-ui/icons/Favorite";
import ShareIcon from "@material-ui/icons/Share";
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";
import clsx from "clsx";
import fern from "../../images/fern.png";
import plantStyles from "./plantStyles";
import { IPlant } from "../../interface/IPlant";

const Plant = ({
  common,
  botanical,
  height,
  characteristics,
  zones,
}: IPlant): JSX.Element => {
  const classes = plantStyles();

  const [expanded, setExpanded] = useState(false);

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const formatRange = (
    range: (string | null)[] | undefined,
    isHeight: boolean
  ): string | null => {
    if (range) {
      return range[0]
        ? range[1]
          ? (!isHeight ? "Zones: " : "") +
            `${range[0]}-${range[1]}` +
            (isHeight ? `"` : "")
          : `${range[0]}"`
        : "";
    }
    return null;
  };

  return (
    <Card className={classes.root}>
      <CardHeader
        className={classes.title}
        avatar={<Avatar className={classes.avatar}>R</Avatar>}
        action={
          <IconButton>
            <MoreVertIcon />
          </IconButton>
        }
        title={common}
        subheader={botanical}
      />
      <CardMedia className={classes.media} image={fern} title={common} />
      <CardContent className={classes.cardContent}>
        <Typography variant="body2" color="textSecondary" component="p">
          {characteristics}
        </Typography>
        <Divider className={classes.divider} />
        <Grid container justify="space-around">
          <Grid item>
            <Typography variant="h6" color="textSecondary" component="p">
              {formatRange(height, true)}
            </Typography>
          </Grid>
          <Grid item>
            <Typography variant="h6" color="textSecondary" component="p">
              {formatRange(zones, false)}
            </Typography>
          </Grid>
        </Grid>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton>
          <FavoriteIcon />
        </IconButton>
        <IconButton>
          <ShareIcon />
        </IconButton>
        <IconButton
          className={clsx(classes.expand, {
            [classes.expandOpen]: expanded,
          })}
          onClick={handleExpandClick}
        >
          <ExpandMoreIcon />
        </IconButton>
      </CardActions>
      <Collapse in={expanded} timeout="auto" unmountOnExit>
        <CardContent>
          <Typography paragraph>{characteristics}</Typography>
        </CardContent>
      </Collapse>
    </Card>
  );
};

export default Plant;
