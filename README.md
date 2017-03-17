# Point Clustering

[![Build Status](https://travis-ci.org/smira/go-point-clustering.svg?branch=master)](https://travis-ci.org/smira/go-point-clustering)
[![Coverage Status](https://coveralls.io/repos/smira/go-point-clustering/badge.svg?branch=master)](https://coveralls.io/r/smira/go-point-clustering?branch=master)
[![GoDoc](https://godoc.org/github.com/smira/go-point-clustering?status.svg)](https://godoc.org/github.com/smira/go-point-clustering)

(Lat, lon) points fast clustering using [DBScan](https://en.wikipedia.org/wiki/DBSCAN) algorithm in Go.

Given set of geo points, this library can find clusters according to specified params. There are several optimizations
applied:

* distance calculation is using "fast" implementations of sine/cosine, with `sqrt` being removed
* to find points within `eps` distance [k-d tree](https://en.wikipedia.org/wiki/K-d_tree) is being used
* edge case handling of identical points being present in the set

## Usage

Build list of points:

```go
    points := cluster.PointList{{30.258387, 59.951557}, {30.434124, 60.029499}, ...}
```

Pick settings for DBScan algorithm:

* `eps` is clustering radius (in kilometers)
* `minPoints` is number of points in `eps`-radius of base point to consider it being part of the cluster

`eps` and `minPoints` together define minimum density of the cluster.

Run DBScan:

```go
    clusters, noise := cluster.DBScan(points, 0.8, 10) // eps is 800m, 10 points minimum in eps-neighborhood
```

`DBScan` function returns list of clusters (each `Cluster` being reference to the list of source `points`) and list
of point indexes which don't fit into any cluster (`noise`).
