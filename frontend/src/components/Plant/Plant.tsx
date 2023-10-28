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
import { addUserFavorite, removeUserFavorite } from "../../api/users";

type PlantProps = IPlant & {userID: string};

const Plant = ({ id, common, botanical, height, characteristics, zones, favorited, userID }: PlantProps): JSX.Element => {
  const classes = plantStyles();

  const [expanded, setExpanded] = useState(false);
  const [favorite, setFavorite] = useState(favorited);

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const handleAddFavorite = async () => {
    try {
      const response = await addUserFavorite({ userID: userID, plantID: id });
      if (response.data.success) {
        setFavorite(true);
      }
    } catch (response) {
      console.log(response);
    }
  };

  const handleRemoveFavorite = async () => {
    try {
      const response = await removeUserFavorite({ userID: userID, plantID: id });
      if (response.data.success) {
        setFavorite(false);
      }
    } catch (response) {
      console.log(response);
    }
  };

  const formatRange = (range: (string | null)[] | undefined, isHeight: boolean): string | null => {
    if (range) {
      return range[0]
        ? range[1]
          ? (!isHeight ? "Zones: " : "") + `${range[0]}-${range[1]}` + (isHeight ? `"` : "")
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
        <IconButton
          onClick={() => {
            setFavorite(!favorite);
            favorite ? handleRemoveFavorite() : handleAddFavorite();
          }}
        >
          <FavoriteIcon color={favorite ? "primary" : "disabled"} />
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
