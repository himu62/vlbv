using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices.WindowsRuntime;
using Windows.ApplicationModel.DataTransfer;
using Windows.Foundation;
using Windows.Foundation.Collections;
using Windows.Storage;
using Windows.UI.Xaml;
using Windows.UI.Xaml.Controls;
using Windows.UI.Xaml.Controls.Primitives;
using Windows.UI.Xaml.Data;
using Windows.UI.Xaml.Input;
using Windows.UI.Xaml.Media;
using Windows.UI.Xaml.Navigation;

// 空白ページの項目テンプレートについては、https://go.microsoft.com/fwlink/?LinkId=402352&clcid=0x411 を参照してください

namespace VLBV
{
    /// <summary>
    /// それ自体で使用できる空白ページまたはフレーム内に移動できる空白ページ。
    /// </summary>
    public sealed partial class MainPage : Page
    {
        public ObservableCollection<ListItem> Items { get; } = new ObservableCollection<ListItem>();
        
        public MainPage()
        {
            this.InitializeComponent();
        }

        private async void Page_Drog(object sender, DragEventArgs e)
        {
            var d = e.GetDeferral();
            var folders = await e.DataView.GetStorageItemsAsync();

            foreach (var folder in folders)
            {
                if (folder is IStorageFolder)
                {
                    var item = new ListItem()
                    {
                        Path = folder.Path,
                        IsRecursive = true
                    };
                    Items.Add(item);
                }
            }

            d.Complete();
        }

        private async void Page_DragOver(object sender, DragEventArgs e)
        {
            var d = e.GetDeferral();
            try
            {
                var items = await e.DataView.GetStorageItemsAsync();
                if (items != null)
                {
                    e.AcceptedOperation = DataPackageOperation.Copy;
                }
            }
            catch (Exception exp)
            {
                throw exp;
            }
            finally
            {
                d.Complete();
            }
        }
    }
}
