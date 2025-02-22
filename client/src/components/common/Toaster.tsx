import { Alert } from "@chakra-ui/react";
import { FC, ReactElement } from "react";

const Toaster: FC = (): ReactElement => {
  return (
    <Alert.Root
      top={"15%"}
      left={"50%"}
      width={"90%"}
      pos={"fixed"}
      zIndex={"1000"}
      status={"success"}
      transform={"translate(-50%, -50%)"}
    >
      <Alert.Indicator />
      <Alert.Title>Done, Senpai :3</Alert.Title>
    </Alert.Root>
  );
}

export default Toaster;