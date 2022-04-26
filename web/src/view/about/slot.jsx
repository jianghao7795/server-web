const Slot = (props, { slots }) => {
  console.log(props);
  return <div>{slots.default?.()}</div>;
};

export default Slot;
