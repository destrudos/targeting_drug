import React from 'react';

import { Metric } from 'types';

import BadgeTag from './BadgeTag';

interface Props {
  metric: Metric;
}

const MetricBadgeTag: React.FC<Props> = ({ metric }: Props) => {
  return (
    <BadgeTag label={metric.name} tooltip={metric.type}>
      {metric.type.substring(0, 1).toUpperCase()}
    </BadgeTag>
  );
};

export default MetricBadgeTag;
