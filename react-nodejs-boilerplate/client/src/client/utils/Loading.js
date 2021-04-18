import { StageSpinner } from "react-spinners-kit";

const Loading = () => {
  return (
    <div
      style={{ display: "flex", alignItems: "center", justifyContent: "center", height: "100vh" }}
    >
      <StageSpinner size={60} color="#686769" loading={true} />
    </div>
  );
};

export default Loading;
