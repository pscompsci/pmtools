import csv
import comments
import credits
import gar
import revenue
import roles
import sys, getopt

rows = 100

def mock_data(n):
    r = revenue.mock_revenue(n)
    c = credits.mock_credits(r)
    g = gar.mock_gar(r)
    p = roles.mock_roles(r)
    d = comments.mock_comments(r, 200)
    return r, c, g, p, d

def save_csv(titles, data, filename):
    with open(filename, 'w') as f:
        writer = csv.writer(f, delimiter=',', quotechar='"',
                            quoting=csv.QUOTE_MINIMAL)
        writer.writerow(titles)
        writer.writerows(data)

def main():
    global rows

    if len(sys.argv) > 1:
        try:
            rows = int(sys.argv[1])
        except Exception:
            print('main.py <number of mock row>')
            sys.exit(2)

    r, c, g, p, d = mock_data(rows)
    save_csv(revenue.titles, r, "ssrs.csv")
    save_csv(credits.titles, c, "credits.csv")
    save_csv(gar.titles, g, "gar.csv")
    save_csv(roles.titles, p, "roles.csv")
    save_csv(comments.titles, d, "comments.csv")

if __name__ == '__main__':
    main()
