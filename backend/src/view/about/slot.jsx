const Slot = (props, { slots }) => {
  // console.log(props, slots);
  const { action = [] } = props;
  return <div>{action.map((i) => slots[i]?.())}</div>;
};

export default Slot;
