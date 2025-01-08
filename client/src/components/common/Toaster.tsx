import { Alert } from "@chakra-ui/react";

export default function Toaster() {
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
      <Alert.Title>Added to the Library</Alert.Title>
    </Alert.Root>
  );
}
