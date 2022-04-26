const Slot = (props, { slots }) => {
  // console.log(props);
  const { action = [] } = props;
  return <div>{action.map((i) => slots[i]?.())}</div>;
};

export default Slot;
