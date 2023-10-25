import { Grid } from "@material-ui/core";
import { useState, useEffect } from "react";
import { v4 as uuidv4 } from "uuid";
import { IPlant } from "../../interface/IPlant";
import { getAllPlants } from "../../api/plants";
import Plant from "../Plant/Plant";
import cardContainerStyles from "./cardContainerStyles";
import { useQuery } from "react-query";

type ContainerProps = {
  searchValue: string;
};

const CardContainer = ({ searchValue }: ContainerProps): JSX.Element => {
  const classes = cardContainerStyles();
  const plantsQuery = useQuery("plants", getAllPlants);
  const [visiblePlants, setVisiblePlants] = useState(plantsQuery.data as IPlant[]);

  useEffect(() => {
    if (plantsQuery.isIdle) {
      return;
    }
    if (plantsQuery.isSuccess) {
      const plants = plantsQuery.data;
      if (plants?.length > 0) {
        setVisiblePlants(plants);
      }
    }
    if (plantsQuery.isError) {
      //TODO: handle errors
    }
  }, [plantsQuery.isIdle, plantsQuery.isSuccess, plantsQuery.isError, plantsQuery.data]);

  return (
    <Grid container spacing={0} className={classes.main}>
      {visiblePlants &&
        visiblePlants
          .filter((plant) => plant && plant.common?.match(RegExp(searchValue, "i")))
          .map((props: IPlant | null) => (
            <Grid item xs={12} sm={6} md={4} lg={3} xl={2}>
              <Plant key={uuidv4()} {...props} />
            </Grid>
          ))}
    </Grid>
  );
};

export default CardContainer;
