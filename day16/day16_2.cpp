#include <bits/stdc++.h>
#include <ext/pb_ds/assoc_container.hpp>
#include <ext/pb_ds/tree_policy.hpp>
using namespace std;
using namespace __gnu_pbds;

template <class T>
using ordered_set = tree<T, null_type, less<T>, rb_tree_tag, tree_order_statistics_node_update>;

#define int long long int
#define pb push_back
#define pi pair<int, int>
#define pii pair<int, pi>
#define fir first
#define sec second
#define MAXN 205
#define mod 998244353

int dist[MAXN][MAXN][4];
int dist2[MAXN][MAXN][4];
bool vis[MAXN][MAXN][4];
bool vis2[MAXN][MAXN][4];
int revdir[] = {2, 3, 0, 1};
int dx[] = {-1, 0, 1, 0};
int dy[] = {0, 1, 0, -1};

signed main()
{
  ios_base::sync_with_stdio(false);
  cin.tie(NULL);
  string ss;
  vector<string> v;
  while (cin >> ss)
  {
    v.pb(ss);
  }
  int n = ss.size();
  pi s, e;
  for (int i = 0; i < n; i++)
  {
    for (int j = 0; j < n; j++)
    {
      if (v[i][j] == 'S')
        s = {i, j};
      if (v[i][j] == 'E')
        e = {i, j};
      for (int k = 0; k < 4; k++)
      {
        dist[i][j][k] = 1e18;
        dist2[i][j][k] = 1e18;
      }
    }
  }
  {
    priority_queue<pair<pi, pi>, vector<pair<pi, pi>>, greater<pair<pi, pi>>> pq;
    pq.push({{0, 1}, s});
    dist[s.fir][s.sec][1] = 0;
    while (!pq.empty())
    {
      auto [x, y] = pq.top().sec;
      auto [d, dir] = pq.top().fir;
      pq.pop();
      if (vis[x][y][dir])
      {
        continue;
      }
      vis[x][y][dir] = 1;
      {
        int nxt = (dir + 1) % 4;
        if (dist[x][y][nxt] > dist[x][y][dir] + 1000 && v[x][y] != '#')
        {
          dist[x][y][nxt] = dist[x][y][dir] + 1000;
          pq.push({{dist[x][y][nxt], nxt}, {x, y}});
        }
      }
      {
        int nxt = (dir - 1 + 4) % 4;
        if (dist[x][y][nxt] > dist[x][y][dir] + 1000 && v[x][y] != '#')
        {
          dist[x][y][nxt] = dist[x][y][dir] + 1000;
          pq.push({{dist[x][y][nxt], nxt}, {x, y}});
        }
      }
      int i = x + dx[dir], j = y + dy[dir];
      if (i >= 0 && i < n && j >= 0 && j < n && dist[i][j][dir] > dist[x][y][dir] + 1 && v[i][j] != '#')
      {
        dist[i][j][dir] = dist[x][y][dir] + 1;
        pq.push({{dist[i][j][dir], dir}, {i, j}});
      }
    }
  }
  {
    priority_queue<pair<pi, pi>, vector<pair<pi, pi>>, greater<pair<pi, pi>>> pq;
    for (int i = 0; i < 4; i++)
    {
      pq.push({{0, i}, e});
      dist2[e.fir][e.sec][i] = 0;
    }
    while (!pq.empty())
    {
      auto [x, y] = pq.top().sec;
      auto [d, dir] = pq.top().fir;
      pq.pop();
      if (vis2[x][y][dir])
      {
        continue;
      }
      vis2[x][y][dir] = 1;
      {
        int nxt = (dir + 1) % 4;
        if (dist2[x][y][nxt] > dist2[x][y][dir] + 1000 && v[x][y] != '#')
        {
          dist2[x][y][nxt] = dist2[x][y][dir] + 1000;
          pq.push({{dist2[x][y][nxt], nxt}, {x, y}});
        }
      }
      {
        int nxt = (dir - 1 + 4) % 4;
        if (dist2[x][y][nxt] > dist2[x][y][dir] + 1000 && v[x][y] != '#')
        {
          dist2[x][y][nxt] = dist2[x][y][dir] + 1000;
          pq.push({{dist2[x][y][nxt], nxt}, {x, y}});
        }
      }
      int i = x + dx[dir], j = y + dy[dir];
      if (i >= 0 && i < n && j >= 0 && j < n && dist2[i][j][dir] > dist2[x][y][dir] + 1 && v[i][j] != '#')
      {
        dist2[i][j][dir] = dist2[x][y][dir] + 1;
        pq.push({{dist2[i][j][dir], dir}, {i, j}});
      }
    }
  }
  int ans = 1e18;
  for (int dir = 0; dir < 4; dir++)
  {
    ans = min(ans, dist[e.fir][e.sec][dir]);
  }
  int qt = 0;
  for (int i = 0; i < n; i++)
  {
    for (int j = 0; j < n; j++)
    {
      int mn = 1e18;
      for (int dir = 0; dir < 4; dir++)
        mn = min(mn, dist[i][j][dir] + dist2[i][j][revdir[dir]]);
      if (mn == ans)
        qt++;
    }
  }
  cout << qt << endl;
  return 0;
}