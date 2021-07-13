import { Grid } from "@material-ui/core";
import plants from "../../data/plants.json";
import { v4 as uuidv4 } from "uuid";
import { IPlant } from "../../interface/IPlant";
import Plant from "../Plant/Plant";
import cardContainerStyles from "./cardContainerStyles";

type ContainerProps = {
  searchValue: string;
};

export function isPlant(object: unknown): boolean {
  return (
    Object.prototype.hasOwnProperty.call(object, "common") &&
    Object.prototype.hasOwnProperty.call(object, "botanical") &&
    Object.prototype.hasOwnProperty.call(object, "height") &&
    Object.prototype.hasOwnProperty.call(object, "characteristics") &&
    Object.prototype.hasOwnProperty.call(object, "zones") &&
    Object.prototype.hasOwnProperty.call(object, "benefits")
  );
}

const CardContainer = ({ searchValue }: ContainerProps): JSX.Element => {
  const classes = cardContainerStyles();

  return (
    <Grid container spacing={0} className={classes.main}>
      {plants
        .filter(
          (plant) =>
            plant &&
            isPlant(plant) &&
            plant.common.match(RegExp(searchValue, "i"))
        )
        .map((props: IPlant | null) => (
          <Grid item xs={12} sm={6} md={4} lg={3} xl={2}>
            <Plant key={uuidv4()} {...props} />
          </Grid>
        ))}
    </Grid>
  );
};

export default CardContainer;
