import { Grid } from "@material-ui/core";
import { useState, useEffect } from "react";
import { IPlant } from "../../interface/IPlant";
import { getAllPlantsForUser } from "../../api/plants";
import Plant from "../Plant/Plant";
import cardContainerStyles from "./cardContainerStyles";
import { useQuery } from "react-query";
import { IUser } from "../../interface/IUser";

type ContainerProps = {
  searchValue: string;
  loggedInUser: IUser;
};

const LoggedInUserPlants = ({ searchValue, loggedInUser }: ContainerProps): JSX.Element => {
  const classes = cardContainerStyles();
  const plantsQuery = useQuery('userPlants', () => getAllPlantsForUser({userID: loggedInUser.id}));
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
          .filter((plant) => plant.common?.match(RegExp(searchValue, "i")))
          .map((props: IPlant) => (
            <Grid item xs={12} sm={6} md={4} lg={3} xl={2}>
              <Plant key={props.id} userID={loggedInUser.id} favorited={true} {...props} />
            </Grid>
          ))}
    </Grid>
  );
};

export default LoggedInUserPlants;
